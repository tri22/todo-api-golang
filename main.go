package main

import (
	"GoProject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes
	routes.TaskRoutes(router)

	// Start server
	err := router.Run(":8080")
	if err != nil {
		println("Failed to start server:", err.Error())
		return
	}
}
