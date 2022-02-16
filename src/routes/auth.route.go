package routes

import (
	"tiddly/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes (r *gin.RouterGroup) {
	r.POST("/signin",controllers.Authentication)
}