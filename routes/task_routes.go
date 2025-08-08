package routes

import (
	"GoProject/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {

	tasks := router.Group("/tasks")
	{
		tasks.GET("/", controllers.GetTasks)
		tasks.GET("/:id", controllers.GetTaskByID)
		tasks.POST("/", controllers.CreateTask)
		tasks.PUT("/:id", controllers.UpdateTask)
		tasks.DELETE("/:id", controllers.DeleteTask)
	}
}
