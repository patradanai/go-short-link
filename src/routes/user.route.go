package routes

import (
	"tiddly/src/controllers"
	"tiddly/src/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.PUT("/userinfo", middlewares.Authorization(), middlewares.SaveFileUpload("image"), controllers.UpdateUserInfo)

	r.POST("/oauth", middlewares.Authorization(), controllers.CreateOauthClient)

	r.POST("/oauth/regen", middlewares.Authorization(), controllers.GenOauthClient)

	r.PUT("/oauth/:oathId", middlewares.Authorization(), controllers.RevokeOauthClient)
}
