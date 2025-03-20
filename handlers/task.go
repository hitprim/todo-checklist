package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-api/database"
	"todo-api/models"
)

func GetTasks(c *gin.Context) {
	db := database.Connect()
	var tasks []models.Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Connect()
	result := db.Create(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Successfully added task"})
}

func UpdateTask(c *gin.Context) {
	db := database.Connect()
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Updates(&task)
	c.JSON(http.StatusOK, gin.H{"Updated id": task.ID})
}

func DeleteTask(c *gin.Context) {
	db := database.Connect()
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := task.ID
	db.Unscoped().Delete(&task, id)
	c.JSON(http.StatusOK, gin.H{"Deleted": id})
}
