/*
main.go
v0.1.0
1/2/24

This file initializes the web server
*/
package main

import (
	"context"
	"net/http"
	"os"
	"strconv"

	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"kademlia.io/handlers"
	"kademlia.io/server"
)

const (
	endpoint_prefix = "/api/v1"
)

func init() {
	//Loading .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load .env file %v\n", err)
	}
	// Setting log level
	LOG_LEVEL := (os.Getenv("LOG_LEVEL"))
	if LOG_LEVEL == "" {
		LOG_LEVEL = "DEBUG"
	}
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(server.GetLogLevel(LOG_LEVEL))
}

func main() {
	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Unable to parse PORT environment variable %v\n", err)
	}

	// If this env is set (doesn't matter the value), then we will use true
	ENABLE_CORS_LOGGING := os.Getenv("ENABLE_CORS_LOGGING")

	if ENABLE_CORS_LOGGING != "" {
		log.Debug("CORS logging enabled")
	}

	log.Debug("Starting Service")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:              PORT,
		EnableCorsLogging: ENABLE_CORS_LOGGING != "",
	})

	if err != nil {
		log.Fatalf("Unable to create server %v\n", err)
	}

	router := mux.NewRouter()

	// Bind routes to the router
	bindRoutes(s, router)

	// Wrap the router with the metrics middleware
	wrappedRouter := server.Middleware(router)

	// Start the server with the wrapped router
	s.Start(wrappedRouter)
}

func bindRoutes(s server.Server, r *mux.Router) {
	// Profiling
	r.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)

	// Home Route
	r.HandleFunc(endpoint_prefix+"/version", handlers.GetVersionRoute(s)).Methods("GET")

	// AES Routes
	r.HandleFunc(endpoint_prefix+"/encrypt", handlers.AesEncrypt(s)).Methods("POST")
	r.HandleFunc(endpoint_prefix+"/decrypt", handlers.AesDecrypt(s)).Methods("POST")

	// ECC Routes
	r.HandleFunc(endpoint_prefix+"/ed25519", handlers.Ed25519Keypair(s)).Methods("GET")
	r.HandleFunc(endpoint_prefix+"/ed25519/sign", handlers.Ed25519Signature(s)).Methods("POST")
	r.HandleFunc(endpoint_prefix+"/ed25519/verify", handlers.Ed25519Verification(s)).Methods("POST")
}
