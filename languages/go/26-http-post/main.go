package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://postman-echo.com/post"

	// 1. Define the data you want to send
	data := map[string]string{
		"username": "gopher123",
		"email":    "gopher@golang.org",
	}

	// 2. Convert (marshal) your data into JSON bytes
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// 3. Create an HTTP client with a timeout (Crucial for production!)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 4. Send the POST request
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close() // Always close the body to prevent memory leaks

	// 5. Read and print the response status and body
	log.Printf("Response Status: %s", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	log.Printf("Response Body:\n%s", string(body))
}
