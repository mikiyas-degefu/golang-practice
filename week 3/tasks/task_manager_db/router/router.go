package router

import (
	"task_manager_db/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/tasks")
	{
		api.GET("", controllers.GetTasks)
		api.POST("", controllers.CreateTask)
		api.GET(":id", controllers.GetTask)
		api.PUT(":id", controllers.UpdateTask)
		api.DELETE(":id", controllers.DeleteTask)
	}

	return r
}
