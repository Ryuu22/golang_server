package config

import (
	"errors"
	"os"
)

const (
	// Mysql constants
	DBUser = "DB_USER"
	DBPass = "DB_PASS"
	DBHost = "DB_HOST"
	DBName = "DB_NAME"

	// Server constants
	Port = "8080"

	// JWT constants
	JWTSecret = "JWT_SECRET" // TODO: Change to ENV VAR
	Issuer    = "golang_server.dankbueno.com"
)

func ValidateConfig() error {
	if os.Getenv(DBUser) == "" {
		return errors.New("DB_USER is not set")
	}
	if os.Getenv(DBPass) == "" {
		return errors.New("DB_PASS is not set")
	}
	if os.Getenv(DBHost) == "" {
		return errors.New("DB_HOST is not set")
	}
	if os.Getenv(DBName) == "" {
		return errors.New("DB_NAME is not set")
	}
	return nil
}
