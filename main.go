package main

import (
	"log"
	"net/http"
	"service-catalog/auth"
	"service-catalog/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Authentication Middleware
	r.Use(auth.AuthMiddleware)

	// Service Handlers
	r.HandleFunc("/services", handlers.GetServices).Methods("GET")
	r.HandleFunc("/services/{id}", handlers.GetService).Methods("GET")
	r.HandleFunc("/services/{id}/versions", handlers.GetServiceVersions).Methods("GET")

	// Prometheus Monitoring
	// prometheus.RegisterPrometheus(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
