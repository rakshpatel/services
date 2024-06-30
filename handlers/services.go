package handlers

import (
	"encoding/json"
	"net/http"
	"service-catalog/auth"
	"service-catalog/backend"
	"strconv"

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

	// Get parameters for pagination
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	// Set page, limit to default if not provided
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	// Get name parameter to support filtering
	name := r.URL.Query().Get("name")

	services, err := backend.GetAllServices(page, limit, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(services)
}

func GetService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	service, err := backend.GetServiceByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(service)
}

func GetServiceVersions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	versions, err := backend.GetServiceVersions(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(versions)
}
