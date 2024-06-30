package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"service-catalog/auth"
	"service-catalog/backend"

	"github.com/gorilla/mux"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

func GetServices(w http.ResponseWriter, r *http.Request) {
	log.Println("GetServices called")
	services, err := backend.GetServices(backend.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(services)
}

func GetService(w http.ResponseWriter, r *http.Request) {
	log.Println("GetService called")
	vars := mux.Vars(r)
	id := vars["id"]
	service, err := backend.GetService(backend.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if service == nil {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(service)
}

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
