package routes

import "github.com/gin-gonic/gin"

func LinkRoutes(r *gin.RouterGroup) {
	r.POST("/create-link")

	r.GET("/:referenceId")
}
