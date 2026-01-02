package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Server represents the HTTP server
type Server struct {
	config     *Config
	mux        *http.ServeMux
	httpServer *http.Server
}

// NewServer creates a new server instance
func NewServer(cfg *Config) *Server {
	mux := http.NewServeMux()
	server := &Server{
		config: cfg,
		mux:    mux,
		httpServer: &http.Server{
			Addr:         ":" + strconv.Itoa(cfg.Port),
			Handler:      mux,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}

	// Register routes
	server.registerRoutes()

	return server
}

// Start starts the HTTP server
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// registerRoutes registers all API routes
func (s *Server) registerRoutes() {
	// Health check endpoint
	s.mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"healthy"}`)
	})

	// Root endpoint
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message":"Frontrun Monitor API","version":"1.0.0"}`)
	})
}
