package configs

import (
	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables
func LoadEnv(env string) {
	filename := ".env"
	switch env {
	case "local":
		filename = ".env.local"
	case "staging":
		filename = ".env.staging"
	case "production":
		filename = ".env.production"
	default:
		log.Error("Unknown environment: ", env)
	}

	err := godotenv.Load(filename)
	if err != nil {
		log.Error("Error loading .env file: ", err)
	}
}
