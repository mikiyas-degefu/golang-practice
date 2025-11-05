package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
}

func updateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}

			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Task Updated",
			})
			return
		}
	}
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, val := range tasks {
		if val.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Task Removed",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Task not found",
	})
}

func addTask(c *gin.Context) {
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Task Created",
	})
}

func main() {
	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.POST("/tasks", addTask)

	router.Run("localhost:8000")
}
