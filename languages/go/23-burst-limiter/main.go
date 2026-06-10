package main

import (
	"fmt"
	"net/http"
	"time"
)

// 1. Create the global bucket (rate limiter).
// It has a maximum burst capacity of 3 tokens.
var burstyLimiter = make(chan time.Time, 3)

func main() {
	// 2. Pre-fill the bucket so the server can handle an immediate burst.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 3. Start the steady refill process in the background.
	// It drips 1 new token into the bucket every 500ms.
	go func() {
		for t := range time.Tick(500 * time.Millisecond) {
			select {
			case burstyLimiter <- t:
				// Token added successfully!
			default:
				// The bucket is already full (3 tokens).
				// Drop the token so we don't freeze this background loop.
			}
		}
	}()

	// 4. Define our web endpoint and start the server.
	http.HandleFunc("/api", handleRequest)

	fmt.Println("Server starting on :8080... Press Refresh rapidly in your browser!")
	http.ListenAndServe(":8080", nil)
}

// This function runs automatically inside a brand-new Goroutine
// EVERY time a user hits http://localhost:8080/api
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// PAY THE TOLL: Try to pull a token out of the bucket.
	// If the bucket is empty, this line pauses this specific user's request.
	<-burstyLimiter

	// Once the token is acquired, process the request.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success! Time processed: %s\n", time.Now().Format("15:04:05.000"))
}
