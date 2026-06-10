package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LayeredLimiter struct {
	mu           sync.Mutex
	userCounters map[string]int
	maxPerUser   int
	window       time.Duration

	// Global safety net: The entire server cannot process
	// more than 100 concurrent requests at any given millisecond.
	globalCapacity chan struct{}
}

func main() {
	limiter := &LayeredLimiter{
		userCounters:   make(map[string]int),
		maxPerUser:     5,
		window:         10 * time.Second,
		globalCapacity: make(chan struct{}, 100), // Max 100 global slots
	}

	// (Assume background routines clear the user map and refill the global capacity)

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		userIP := r.RemoteAddr

		// --- CHECKPOINT 1: Global Server Capacity ---
		select {
		case limiter.globalCapacity <- struct{}{}:
			// There is room on the server! Slot taken.
			defer func() { <-limiter.globalCapacity }() // Release slot when done
		default:
			// Server is completely melting down. Reject immediately.
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "Server is busy. Please try again in a moment.\n")
			return
		}

		// --- CHECKPOINT 2: Individual User Behavior ---
		limiter.mu.Lock()
		if limiter.userCounters[userIP] >= limiter.maxPerUser {
			limiter.mu.Unlock()
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, "You have exceeded your individual limit.\n")
			return
		}
		limiter.userCounters[userIP]++
		limiter.mu.Unlock()

		// If they pass both checkpoints, process the request...
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Login successful!\n")
	})

	http.ListenAndServe(":8080", nil)
}
