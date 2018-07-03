package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	route := gin.Default()

	initWebRoute(route)
	initStaticRoute(route)
	return route
}