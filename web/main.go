package main

import (
	"github.com/devopsjester/weather-app-demo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin framework
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	// Set up the routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
