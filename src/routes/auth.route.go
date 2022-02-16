package routes

import (
	"tiddly/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/signup", controllers.RegisterUser)
	r.POST("/signin", controllers.Authentication)
	r.POST("/signout")
	r.POST("/refresh", controllers.RefreshToken)
}
