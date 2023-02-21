package routes

import "github.com/gin-gonic/gin"

func WebRouteInstance(routes *gin.Engine) *gin.Engine {
	routes.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return routes
}
