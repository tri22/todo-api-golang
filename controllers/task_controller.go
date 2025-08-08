package controllers

import (
	"net/http"
	"strconv"

	"GoProject/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := models.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.TaskCreation
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty"})
		return
	}
	id := models.AddTask(task)
	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully", "ID": id})
}

func UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updateData struct {
		Completed bool `json:"completed"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	success := models.UpdateTaskStatus(id, updateData.Completed)
	if !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	success := models.DeleteTask(id)
	if !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.Status(http.StatusOK)
}

func GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	task := models.GetTaskByID(id)
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}
