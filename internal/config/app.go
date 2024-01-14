package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// App is the application settings
type AppConfig struct {
	Url  string
	Port string
	Name string
}

var App AppConfig = AppConfig{
	Url:  os.Getenv("APP_URL"),
	Port: os.Getenv("APP_PORT"),
	Name: os.Getenv("APP_NAME"),
}
