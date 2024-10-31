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
		taskRoutes.PUT("/:taskID", handlers.UpdateTask, middlewares.IDTaskValidateMiddleware())
		taskRoutes.DELETE("/:taskID", handlers.RemoveTask, middlewares.IDTaskValidateMiddleware())
		taskRoutes.POST("/:taskId", handlers.CheckTask)

	}
}
