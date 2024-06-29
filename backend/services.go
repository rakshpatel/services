package backend

import (
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
