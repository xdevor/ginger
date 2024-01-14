package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// Database is the database settings
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

var Database DatabaseConfig = DatabaseConfig{
	Host:     os.Getenv("DB_HOST"),
	Port:     os.Getenv("DB_PORT"),
	User:     os.Getenv("DB_USERNAME"),
	Password: os.Getenv("DB_PASSWORD"),
	Database: os.Getenv("DB_DATABASE"),
}
