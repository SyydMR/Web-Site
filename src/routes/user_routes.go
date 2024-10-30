package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/SyydMR/Web-Site/src/handlers"
)

func addUserRoute(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/users", handlers.GetAllUsers)

		userRoutes.POST("/login", handlers.LoginHandler)

		userRoutes.POST("/register", handlers.RegisterHandler)

		userRoutes.GET("/users/:userId", handlers.GetUserByIDHandler)
	}
}
