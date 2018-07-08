package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yushuailiu/easygin/resources"
)

func initStaticRoute(route *gin.Engine) {
	staticAsserts()

	staticRoute := route.Group("/")
	staticRoute.StaticFS("/resources", resources.StaticAsserts)
}

func staticAsserts() {
	for index, item := range resources.StaticAsserts.Files {
		if item.IsDir() {
			delete(resources.StaticAsserts.Files, index)
		}
	}
}
