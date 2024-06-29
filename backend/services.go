package backend

import (
	"errors"
	"service-catalog/datamodels"
)

var services = []datamodels.Service{
	{ID: "1", Name: "Service A", Description: "Description for Service A", Versions: []string{"v1.0", "v1.1"}},
	{ID: "2", Name: "Service B", Description: "Description for Service B", Versions: []string{"v2.0", "v2.1"}},
}

func GetAllServices() ([]datamodels.Service, error) {
	return services, nil
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
