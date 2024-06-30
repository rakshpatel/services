package backend

import (
	"database/sql"
	"fmt"
	"log"
	"service-catalog/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// DB initialization, connecting to PG DB
// Creates connection string and connects to PG DB else throws error
func InitDB(cfg config.Config) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %q", err)
	}

	log.Println("Successfully connected to the database")
}
