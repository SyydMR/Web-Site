package routes

import (
	"github.com/SyydMR/Web-Site/src/handlers"
	"github.com/SyydMR/Web-Site/src/middlewares"
	"github.com/gin-gonic/gin"
)

func addTaskRoute(r *gin.Engine) {
	taskRoutes := r.Group("/tasks", middlewares.AuthMiddleware())
	{
		taskRoutes.GET("", handlers.GetAllTasks)
		taskRoutes.POST("", handlers.AddTask)
		taskRoutes.PUT("/:TaskId", handlers.UpdateTask)
		taskRoutes.DELETE("/:TaskId", handlers.RemoveTask)
		taskRoutes.POST("/:TaskId", handlers.CheckTask)
	}
}
