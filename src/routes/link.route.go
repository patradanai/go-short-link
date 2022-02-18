package routes

import (
	"tiddly/src/controllers"

	"github.com/gin-gonic/gin"
)

func LinkRoutes(r *gin.RouterGroup) {
	r.POST("/create-link", controllers.CreateLink)

	r.GET("/:shortId", controllers.RedirectLink)
}
