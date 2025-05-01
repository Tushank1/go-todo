package controllers

import (
	"todo/database"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, completed FROM todo")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, todo)
	}
	c.JSON(200, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := database.DB.QueryRow(
		"INSERT INTO todo (title, completed) VALUES ($1, $2) RETURNING id",
		todo.Title, todo.Completed).Scan(&todo.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	_, err := database.DB.Exec("DELETE FROM todo WHERE id = $1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec(
		"UPDATE todo SET title = $1, completed = $2 WHERE id = $3",
		todo.Title, todo.Completed, id,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Todo updated"})

	c.JSON(200, todo)
}
