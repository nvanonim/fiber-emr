package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nvanonim/fiber-emr/pkg/configs"
	"github.com/nvanonim/fiber-emr/pkg/routes"
	"github.com/nvanonim/fiber-emr/pkg/utils"
)

func main() {
	fmt.Println("Starting the server...")
	// Load the environment variables
	utils.LoadEnv()
	// Setup the database
	configs.SetupDB()

	// Register the routes
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Run the server
	r.Run()
}
