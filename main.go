package main

import (
	"GoProject/routes"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	routes.TaskRoutes(router)

	// Start server
	err := router.Run(":8080")
	if err != nil {
		println("Failed to start server:", err.Error())
		return
	}
}
