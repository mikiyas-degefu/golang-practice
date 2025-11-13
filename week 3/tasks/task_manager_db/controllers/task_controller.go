package controllers

import (
	"net/http"
	"strconv"
	"task_manager_db/data"
	"task_manager_db/models"

	"github.com/gin-gonic/gin"
)

// GetTasks handles GET /tasks
func GetTasks(c *gin.Context) {
	list := data.GetAll()
	c.JSON(http.StatusOK, gin.H{"tasks": list})
}

// GetTask handles GET /tasks/:id
func GetTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	task, err := data.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// CreateTask handles POST /tasks
func CreateTask(c *gin.Context) {
	var payload models.Task
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := data.Create(payload)
	c.JSON(http.StatusCreated, created)
}

// UpdateTask handles PUT /tasks/:id
func UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var payload models.Task
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := data.Update(id, payload)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteTask handles DELETE /tasks/:id
func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := data.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
