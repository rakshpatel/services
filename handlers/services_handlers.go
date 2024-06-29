package handlers

import (
	"encoding/json"
	"net/http"
	"service-catalog/backend"

	"github.com/gorilla/mux"
)

func GetServices(w http.ResponseWriter, r *http.Request) {
	services, err := backend.GetAllServices()
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
