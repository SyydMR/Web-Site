package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/SyydMR/Web-Site/src/handlers"
)

func addBlogRoute(r *gin.Engine) {
	blogRoutes := r.Group("/blog")
	{
		blogRoutes.GET("/", handlers.GetAllPosts)
		blogRoutes.GET("/:postId", handlers.GetPostByID)
		blogRoutes.POST("/", handlers.CreatePost)
		blogRoutes.PUT("/:postId", handlers.UpdatePost)
		blogRoutes.DELETE("/:postId", handlers.DeletePost)
	}
}
