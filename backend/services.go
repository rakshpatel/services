package backend

import (
	"database/sql"
	"service-catalog/datamodels"
)

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
