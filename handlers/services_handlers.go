package handlers

import (
	"encoding/json"
	"net/http"
	"service-catalog/backend"
	"strconv"

	"github.com/gorilla/mux"
)

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

	services, err := backend.GetAllServices(page, limit)
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
