package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nvanonim/fiber-emr/app/configs"
	"github.com/nvanonim/fiber-emr/app/middlewares"
	"github.com/nvanonim/fiber-emr/app/routes"
)

func main() {
	fmt.Println("Starting the server...")

	initConfigs()

	// Register the routes
	gin.ForceConsoleColor()
	r := gin.Default()
	r.Use(middlewares.RequestLogger())
	r.Use(middlewares.ResponseLogger())
	routes.RegisterRoutes(r)

	// Run the server
	r.Run()
}

func initConfigs() {
	// Setup the logger
	configs.SetupLogger()

	// Check if an environment argument is provided
	var env string
	if len(os.Args) > 1 {
		env = os.Args[1]
	} else {
		env = "local" // Default to "local" if no argument provided
	}

	// Load the environment variables
	configs.LoadEnv(env)
	// Setup the database
	configs.SetupDB()
}
