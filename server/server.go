/*
server.go
v0.1.0
1/2/24

This file defines the underlying web server class. It's used by main.go
*/
package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"kademlia.io/config"
)

// Config
type Config struct {
	Port              int
	EnableCorsLogging bool
}

// Server - server interface to handle connections
type Server interface {
	Config() *Config
}

// Broker - allows to register a new server
type Broker struct {
	config *Config
	router *mux.Router
}

// NewBroker - creates a new broker
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == 0 {
		return nil, errors.New("port must be set")
	}
	broker := &Broker{
		config: config,
	}
	return broker, nil
}

// Config - returns the broker
func (b *Broker) Config() *Config {
	return b.config
}

// Start
func (b *Broker) Start(handler http.Handler) {
	config := config.Get()
	corsHandler := cors.New(cors.Options{
		Debug: b.Config().EnableCorsLogging,
		AllowedHeaders: []string{
			"*",
		},
	})
	handler = corsHandler.Handler(handler)

	port := fmt.Sprintf(":%d", b.Config().Port)

	server := &http.Server{
		Addr:    port,
		Handler: handler,
	}

	// Channel to listen for os signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// WaitGroup to wait for all connections to be closed
	var wg sync.WaitGroup

	log.Printf("Starting Service on port %s and version %s\n", port, config.Version)

	// Goroutine to run the server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	// Block and wait for an interrupt signal
	<-stop
	log.Println("Shutting down the server...")

	// Close connections and wait for them to finish
	wg.Wait()

	log.Println("Server gracefully stopped")
}
