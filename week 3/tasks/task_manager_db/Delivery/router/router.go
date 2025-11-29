package router

import (
	"task_manager_db/Delivery/controllers"
	middleware "task_manager_db/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// --- Public routes ---
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// --- Protected routes (any authenticated user) ---
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	// Tasks accessible by all users
	auth.GET("/tasks", controllers.GetAllTasks)
	auth.GET("/tasks/:id", controllers.GetTaskByID)

	// --- Admin-only routes ---
	admin := auth.Group("/")
	admin.Use(middleware.AdminOnly())

	admin.POST("/tasks", controllers.CreateTask)
	admin.PUT("/tasks/:id", controllers.UpdateTask)
	admin.DELETE("/tasks/:id", controllers.DeleteTask)

	admin.PUT("/promote/:id", controllers.PromoteUser)

	return r
}
