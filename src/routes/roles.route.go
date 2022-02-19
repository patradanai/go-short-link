package routes

import (
	"tiddly/src/controllers"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.RoleList)
	r.POST("/", controllers.CreateRole)
	r.PUT("/:roleId", controllers.UpdateRole)
	r.DELETE("/:roleId", controllers.DeleteRole)
}
