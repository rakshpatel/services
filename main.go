package main

import (
	"log"
	"net/http"
	"service-catalog/auth"
	"service-catalog/backend"
	"service-catalog/config"
	"service-catalog/handlers"

	"github.com/gorilla/mux"
)

func main() {
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
}
