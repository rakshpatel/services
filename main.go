package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-catalog/auth"
	"service-catalog/backend"
	"service-catalog/config"
	"service-catalog/handlers"
	"service-catalog/logger"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	logger.Log.Info("Application starting...")

	setupCloseHandler()

	dbcfg := config.LoadConfig()
	backend.InitDB(dbcfg)

	r := mux.NewRouter()

	// AUthentication endpoint
	r.HandleFunc("/v1/login", handlers.Login).Methods("POST")

	// Service Handlers
	api := r.PathPrefix("/v1/services").Subrouter()
	api.Use(auth.JWTAuth)

	api.HandleFunc("", handlers.GetServices).Methods("GET")
	api.HandleFunc("/{id}", handlers.GetService).Methods("GET")
	api.HandleFunc("/{id}/versions", handlers.GetServiceVersionsDB).Methods("GET")

	// Prometheus Monitoring
	// prometheus.RegisterPrometheus(r)

	// Start the server
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	logger.Log.Info("Application started successfully")
}

func setupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		logger.Log.Info("Received signal:", sig)
		logger.Log.Info("Shutting down...")
		// Perform any cleanup tasks
		logger.Log.Info("Application stopped")
		os.Exit(0)
	}()
}
