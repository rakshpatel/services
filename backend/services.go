package backend

import (
	"database/sql"
	"fmt"
	"service-catalog/datamodels"
	"service-catalog/logger"

	"github.com/sirupsen/logrus"
)

// GetServices gets all the services from the PG DB
func GetServices(db *sql.DB) ([]datamodels.Service, error) {
	logger.Log.Info("Fetching all services from the database")
	rows, err := db.Query("SELECT id, name, description FROM services")
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error fetching all services from database")
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

// GetService fetches a single service from DB
func GetService(db *sql.DB, id string) (*datamodels.Service, error) {
	logger.Log.Info("Fetching a service from the database")
	var service datamodels.Service
	err := db.QueryRow("SELECT id, name, description FROM services WHERE id = $1", id).Scan(&service.ID, &service.Name, &service.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Log.WithFields(logrus.Fields{
				"service_id": id,
			}).Info("No service found in database with the given ID")
			return nil, fmt.Errorf("no service found for ID %s", id)
		}
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error fetching a service from database")
		return nil, err
	}
	return &service, nil
}

// GetServiceVersionsDB, fetches version related data for given service
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
