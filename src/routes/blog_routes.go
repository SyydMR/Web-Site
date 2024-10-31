package routes

import (
	"github.com/SyydMR/Web-Site/src/handlers"
	"github.com/SyydMR/Web-Site/src/middlewares"
	"github.com/gin-gonic/gin"
)

func addBlogRoute(r *gin.Engine) {
	blogRoutes := r.Group("/blogs")
	{
		blogRoutes.GET("/all-posts", handlers.HandlerGetAllPosts)

		blogRoutes.GET("/:postId", handlers.HandlerGetPostByID)

		blogRoutes.POST("/create-post", handlers.HandlerCreateEmptyPost, middlewares.AuthMiddleware()) // redirect to /:postId/update

		updateRoutes := blogRoutes.Group("/:postId/update", middlewares.AuthMiddleware())
		{
			updateRoutes.GET("/", handlers.HandlerGetAllContent)

			updateRoutes.POST("/create-content", handlers.HandlerCreateContent)
			updateRoutes.DELETE("/remove-content/:contentID", handlers.HandlerRemoveContent)

		}



		blogRoutes.GET("/user/all-posts", handlers.HandlerGetUserAllPosts, middlewares.AuthMiddleware())

		// blogRoutes.PUT("/:postId", handlers.UpdatePost)
		blogRoutes.DELETE("/user/:postId/delete-post", handlers.HandlerDeletePost, middlewares.AuthMiddleware())
	}
}
