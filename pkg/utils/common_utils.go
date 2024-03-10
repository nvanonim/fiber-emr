package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables
func LoadEnv() {
	err := godotenv.Load(".env", ".env.local", ".env.production", ".env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv gets the environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}

// GetEnvWithDefault gets the environment variable with a default value
func GetEnvWithDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
