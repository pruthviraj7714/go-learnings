package main

import (
	"day-02/middlewares"
	"day-02/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var nextId = 1

var todos = []models.Todo{}

func addTodoHandler(c *gin.Context) {
	var in models.Todo

	if err := c.ShouldBindBodyWithJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	in.Id = nextId
	nextId++

	todos = append(todos, in)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo Created",
		"data":    in,
	})
}

func updateTodoStatusHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Id",
		})
		return
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos[i].Completed = true
			c.JSON(http.StatusOK, gin.H{
				"message": "Status Successfully Updated",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Todo with give id not found"})
}

func deleteTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Id",
		})
		return
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "todo successfully deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Todo with given id not found"})
}

func getTodosHandler(c *gin.Context) {
	if len(todos) > 0 {
		c.JSON(http.StatusOK, todos)
		return
	} else {
		c.JSON(http.StatusOK, []models.Todo{})
		return
	}
}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	api.POST("/add-todo", addTodoHandler)
	api.PUT("/update-todo/:id", middlewares.AuthMiddleware(), updateTodoStatusHandler)
	api.DELETE("/delete-todo/:id", middlewares.AuthMiddleware(), deleteTodoHandler)
	api.GET("/todos", getTodosHandler)

	router.Run()
}
