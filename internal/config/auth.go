package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// Auth is the auth settings
type AuthConfig struct {
	JWT JWTConfig
}

type JWTConfig struct {
	Secret string
	TTL    string
}

var Auth AuthConfig = AuthConfig{
	JWT: JWTConfig{
		Secret: os.Getenv("JWT_SECRET"),
		TTL:    os.Getenv("JWT_TTL"),
	},
}
