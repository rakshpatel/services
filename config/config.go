package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config, defines DB construct
type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

// Loding config to Config struct for DB
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, falling back to environment variables")
	}

	config := Config{
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "service_catalog"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
	}

	return config
}

// Fetch given key
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
