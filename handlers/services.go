package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"service-catalog/auth"
	"service-catalog/backend"
	"service-catalog/logger"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login, mocks the authentication and provides mock JWT
func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate credentials (this example uses hardcoded values for simplicity)
	if creds.Username == "user" && creds.Password == "password" {
		token, err := auth.GenerateJWT(creds.Username)
		if err != nil {
			http.Error(w, "failed to generate token", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	}
}

// GetServices is handler when /v1/services API is called and calls DB function to get the data
func GetServices(w http.ResponseWriter, r *http.Request) {
	logger.Log.WithFields(logrus.Fields{
		"method":   r.Method,
		"endpoint": "/services",
	}).Info("Handling the get services request")
	services, err := backend.GetServices(backend.DB)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to fetch services")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(services)
	logger.Log.Info("Services retrieved successfully")
}

// GetServices is handler when /v1/service/{ID} API is called and calls DB function to get the data
func GetService(w http.ResponseWriter, r *http.Request) {
	logger.Log.WithFields(logrus.Fields{
		"method":   r.Method,
		"endpoint": "/service",
	}).Info("Handling the get services request")
	vars := mux.Vars(r)
	id := vars["id"]
	service, err := backend.GetService(backend.DB, id)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to fetch service")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if service == nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Service not found")
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(service)
	logger.Log.Info("Service retrieved successfully")
}

// GetServiceVersionsDB is handler when /v1/services/{id}/versions API is called and calls DB function to get the data
func GetServiceVersionsDB(w http.ResponseWriter, r *http.Request) {
	log.Println("GetServiceVersions called")
	vars := mux.Vars(r)
	id := vars["id"]
	versions, err := backend.GetServiceVersionsDB(backend.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(versions)
}
