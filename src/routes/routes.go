package routes

import (
	"github.com/gin-gonic/gin"
	// "net/http"
)

func routeInit() *gin.Engine {
	router := gin.Default()
	addUserRoute(router)
	addTaskRoute(router)
	addBlogRoute(router)
	return router
}

func GetRoute() *gin.Engine {
	router := routeInit()
	return router
}
