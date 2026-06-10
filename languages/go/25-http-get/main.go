package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. Configuration (often read from environment variables)
	port := getEnv("PORT", "8080")
	host := getEnv("HOST", "localhost")
	addr := fmt.Sprintf("%s:%s", host, port)

	// 2. Setup Structured Logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// 3. Explicitly configure the server
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// 4. Handle Graceful Shutdown in a goroutine
	shutdownError := make(chan error, 1)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit // Block until we get a signal (like Ctrl+C)

		logger.Info("server is shutting down...")

		// Give active connections 15 seconds to finish up
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()

	// 5. Start the server (Before blocking, we log!)
	logger.Info("server started", slog.String("host", host), slog.String("port", port))

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server failed to start", "error", err)
		os.Exit(1)
	}

	// Wait for shutdown to complete cleanly
	if err := <-shutdownError; err != nil {
		logger.Error("graceful shutdown failed", "error", err)
	}
	logger.Info("server stopped completely")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
