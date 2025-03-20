package main

import (
	"github.com/gin-gonic/gin"
	"todo-api/database"
	"todo-api/handlers"
)

func main() {
	database.Connect()
	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)
	r.POST("/add", handlers.CreateTask)
	r.POST("/update", handlers.UpdateTask)
	r.DELETE("/delete", handlers.DeleteTask)

	r.Run(":8080")
}
