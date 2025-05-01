package main

import (
	"todo/controllers"
	"todo/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Connect()

	r.GET("/todos", controllers.GetTodo)
	r.POST("/todos", controllers.CreateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)

	r.Run(":8080")
}
