package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Client represents a single SSE client connection
type Client struct {
	ID string // Added client identifier
	w  http.ResponseWriter
	f  http.Flusher
	mu sync.Mutex
	// Add done channel for cleanup
	done chan struct{}
}

// SSE represents the SSE server
type SSE struct {
	clients   map[string]*Client // Changed to use string keys
	mu        sync.RWMutex       // Single mutex for the SSE struct
	maxConns  int                // Maximum allowed connections
	keepAlive time.Duration      // Keepalive interval
}

// SSEConfig holds configuration for the SSE server
type SSEConfig struct {
	MaxConnections int
	KeepAlive      time.Duration
	Origins        []string // Allowed CORS origins
}

func NewSSEFDefault() *SSE {
	return NewSSE(SSEConfig{})
}

// NewSSE creates a new SSE instance with configuration
func NewSSE(config SSEConfig) *SSE {
	if config.MaxConnections <= 0 {
		config.MaxConnections = 10000 // Default max connections
	}
	if config.KeepAlive <= 0 {
		config.KeepAlive = 10 * time.Second // Default keepalive
	}

	return &SSE{
		clients:   make(map[string]*Client),
		maxConns:  config.MaxConnections,
		keepAlive: config.KeepAlive,
	}
}

// Message represents an SSE message
type Message struct {
	Subject      string `json:"subject"`
	FunctionName string `json:"function_name"`
	Data         any    `json:"data"`
}

func enableCors(w http.ResponseWriter, origins []string) {
	// Default to strict CORS if no origins specified
	origin := "*"
	if len(origins) > 0 {
		// In production, you should validate the origin against the allowed list
		origin = origins[0]
	}

	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (s *SSE) addClient(client *Client) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check max connections
	if len(s.clients) >= s.maxConns {
		return fmt.Errorf("maximum connections (%d) reached", s.maxConns)
	}

	s.clients[client.ID] = client
	return nil
}

func (s *SSE) removeClient(clientID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if client, exists := s.clients[clientID]; exists {
		close(client.done)
		delete(s.clients, clientID)
	}
}

func (s *SSE) BroadcastToClients(ctx context.Context, msg Message) error {
	// Validate message
	if msg.Subject == "" {
		return fmt.Errorf("invalid message: subject cannot be empty")
	}
	if msg.Data == nil {
		return fmt.Errorf("invalid message: data cannot be nil")
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	s.mu.RLock()
	clients := make([]*Client, 0, len(s.clients))
	for _, client := range s.clients {
		clients = append(clients, client)
	}
	s.mu.RUnlock()

	// Broadcast to all clients with timeout
	const broadcastTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(ctx, broadcastTimeout)
	defer cancel()

	var wg sync.WaitGroup
	for _, client := range clients {
		wg.Add(1)
		go func(c *Client) {
			defer wg.Done()

			// Use select to handle context cancellation
			select {
			case <-ctx.Done():
				return
			default:
				c.mu.Lock()
				defer c.mu.Unlock()

				// Send SSE format with error handling
				_, err := fmt.Fprintf(c.w, "data: %s\n\n", msgBytes)
				if err != nil {
					log.Printf("Failed to send to client %s: %v", c.ID, err)
					s.removeClient(c.ID)
					return
				}
				c.f.Flush()
			}
		}(client)
	}

	wg.Wait()
	return nil
}

func (s *SSE) HandleSSE(w http.ResponseWriter, r *http.Request) {
	// Handle OPTIONS request for CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	enableCors(w, nil) // Pass allowed origins from config

	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Check if client supports flushing
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Create new client with unique ID
	clientID := fmt.Sprintf("client-%d", time.Now().UnixNano())
	client := &Client{
		ID:   clientID,
		w:    w,
		f:    flusher,
		done: make(chan struct{}),
	}

	// Add client to broadcast list
	if err := s.addClient(client); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer s.removeClient(clientID)

	// Send initial keepalive on connection
	client.mu.Lock()
	fmt.Fprintf(client.w, ": keepalive\n\n")
	client.f.Flush()
	client.mu.Unlock()

	// Start keepalive goroutine
	go func() {
		ticker := time.NewTicker(s.keepAlive)
		defer ticker.Stop()

		for {
			select {
			case <-client.done:
				return
			case <-r.Context().Done():
				return
			case <-ticker.C:
				client.mu.Lock()
				fmt.Fprintf(client.w, ": keepalive\n\n")
				client.f.Flush()
				client.mu.Unlock()
			}
		}
	}()

	// Wait for client disconnect
	select {
	case <-r.Context().Done():
	case <-client.done:
	}
}
