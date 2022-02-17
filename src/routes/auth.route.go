package routes

import (
	"tiddly/src/controllers"
	"tiddly/src/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/signup", controllers.RegisterUser)
	r.POST("/signin", middlewares.Authentication())
	r.POST("/signout")
	r.POST("/refresh", controllers.RefreshToken)
}
