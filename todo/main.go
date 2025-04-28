package main

import (
	"todo/controllers"
	"todo/database"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Connect()
	database.DB.AutoMigrate(&models.Todo{})

	r.GET("/todos", controllers.GetTodo)
	r.POST("/todos", controllers.CreateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	r.Run()
}
