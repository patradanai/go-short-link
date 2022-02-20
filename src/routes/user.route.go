package routes

import (
	"tiddly/src/controllers"
	"tiddly/src/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.PUT("/userinfo", middlewares.Authorization(), middlewares.SaveFileUpload("image"), controllers.UpdateUserInfo)
}
