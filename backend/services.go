package backend

import (
	"database/sql"
	"errors"
	"service-catalog/datamodels"
	"strings"
)

var services = []datamodels.Service{
	{ID: "1", Name: "ServiceA", Description: "Description for Service A", Versions: []string{"v1.0", "v1.1"}},
	{ID: "2", Name: "ServiceB", Description: "Description for Service B", Versions: []string{"v2.0", "v2.1"}},
}

func GetAllServices(page, limit int, filName string) ([]datamodels.Service, error) {
	start := (page - 1) * limit
	end := start + limit

	if start >= len(services) {
		return []datamodels.Service{}, nil
	}

	if end > len(services) {
		end = len(services)
	}

	filteredService := []datamodels.Service{}
	for _, service := range services {
		if filName == "" || strings.Contains(strings.ToLower(service.Name), strings.ToLower(filName)) {
			filteredService = append(filteredService, service)
		}
	}

	if start >= len(filteredService) {
		return []datamodels.Service{}, nil
	}

	if end > len(filteredService) {
		end = len(filteredService)
	}

	return filteredService[start:end], nil
}

func GetServiceByID(id string) (datamodels.Service, error) {
	for _, service := range services {
		if service.ID == id {
			return service, nil
		}
	}
	return datamodels.Service{}, errors.New("service not found")
}

func GetServiceVersions(id string) ([]string, error) {
	for _, service := range services {
		if service.ID == id {
			return service.Versions, nil
		}
	}
	return nil, errors.New("service not found")
}

// DB
func GetServices(db *sql.DB) ([]datamodels.Service, error) {
	rows, err := db.Query("SELECT id, name, description FROM services")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []datamodels.Service
	for rows.Next() {
		var service datamodels.Service
		if err := rows.Scan(&service.ID, &service.Name, &service.Description); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return services, nil
}

func GetService(db *sql.DB, id string) (*datamodels.Service, error) {
	var service datamodels.Service
	err := db.QueryRow("SELECT id, name, description FROM services WHERE id = $1", id).Scan(&service.ID, &service.Name, &service.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &service, nil
}

func GetServiceVersionsDB(db *sql.DB, id string) ([]string, error) {
	rows, err := db.Query("SELECT version FROM service_versions WHERE service_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var versions []string
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return versions, nil
}
