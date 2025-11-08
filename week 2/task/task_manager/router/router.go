package router

import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
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
