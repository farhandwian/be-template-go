package helper

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"
)

func StartServerWithGracefullyShutdown(mux *http.ServeMux) {

	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: cors.AllowAll().Handler(mux),
	}

	// Channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	// Start the server
	go func() {
		fmt.Printf("Server is starting on port %s\n", port)
		serverErrors <- server.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking select
	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s\n", err)
		}
	case <-shutdown:
		fmt.Println("Starting shutdown...")

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		err := server.Shutdown(ctx)
		if err != nil {
			fmt.Printf("Error shutting down server: %s\n", err)
			// Failure/timeout shutting down the server gracefully
			// Forcefully close the server
			err = server.Close()
			if err != nil {
				fmt.Printf("Error closing server: %s\n", err)
			}
		}

		fmt.Println("Server gracefully stopped")
	}

}
