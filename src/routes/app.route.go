package routes

import (
	"tiddly/src/controllers"
	"tiddly/src/middlewares"

	"github.com/gin-gonic/gin"
)

func AppRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.ListAppPackage)

	r.GET("/:packageId", controllers.GetAppPackage)

	r.POST("/", middlewares.Authorization(), middlewares.Permission("ADMINISTRATOR"), controllers.CreateAppPackage)

	r.PUT("/")

	r.DELETE("/:packageId", middlewares.Authorization(), middlewares.Permission("ADMINISTRATOR"), controllers.DeleteAppPackage)
}
