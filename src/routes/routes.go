package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/SyydMR/Web-Site/src/middlewares"

	// "net/http"
)

func routeInit() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.10.1"})
	router.Use(middlewares.Cors())
	addUserRoute(router)
	addTaskRoute(router)
	addBlogRoute(router)
	return router
}

func GetRoute() *gin.Engine {
	router := routeInit()
	return router
}
