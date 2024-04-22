package configs

import (
	"fmt"

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
		fmt.Println("Invalid environment argument. Defaulting to local environment.")
	}
	fmt.Println("Loading environment variables from: ", filename)

	err := godotenv.Load(filename)
	if err != nil {
		fmt.Println("Error loading environment variables: ", err)
	}
}
