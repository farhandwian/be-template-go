package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SendMessage(message string, phoneNumbers ...string) error {
	// URL for sending the message
	// url := "http://localhost:8080/send-message"

	// Create the request body
	requestBody, err := json.Marshal(map[string]any{
		"phone_number": phoneNumbers,
		"message":      message,
	})
	if err != nil {
		return fmt.Errorf("error creating request body: %v", err)
	}

	fmt.Printf(">> %v\n", string(requestBody))

	url := os.Getenv("WHATSAPP_API_URL")

	// Create a new POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK (200)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Printf("Response: %s\n", string(body))
	return nil
}
