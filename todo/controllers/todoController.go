package controllers

import (
	"todo/database"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(200, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&todo)
	c.JSON(200, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Todo{}, "id = ?", id)
	c.Status(204)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := database.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update Todo"})
		return
	}

	c.JSON(200, todo)
}
