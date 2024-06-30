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

	// AUthentication endpoint
	r.HandleFunc("/v1/login", handlers.Login).Methods("POST")

	// Authentication Middleware
	// r.Use(auth.AuthMiddleware)

	// Service Handlers
	api := r.PathPrefix("/v1/services").Subrouter()
	api.Use(auth.JWTAuth)
	// api.Use(auth.AuthMiddleware)

	api.HandleFunc("", handlers.GetServices).Methods("GET")
	api.HandleFunc("/{id}", handlers.GetService).Methods("GET")
	api.HandleFunc("/{id}/versions", handlers.GetServiceVersions).Methods("GET")

	// Prometheus Monitoring
	// prometheus.RegisterPrometheus(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
