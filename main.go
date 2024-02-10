package main

import (
	"github.com/esleirter/e-learning-api/routes"

	"github.com/esleirter/e-learning-api/config"
	"github.com/esleirter/e-learning-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configurations
	config.LoadConfig()

	// Initialize database connection
	utils.InitDB()

	// Initialize Gin engine
	router := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
