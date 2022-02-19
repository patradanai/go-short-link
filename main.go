package main

import (
	"fmt"
	"net/http"
	"tiddly/src/configs"
	"tiddly/src/middlewares"
	"tiddly/src/routes"

	"github.com/gin-gonic/gin"
)

func initialRouter(c *gin.Engine) {
	c.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is Root")
	})

	v1 := c.Group("/api/v1")
	routes.AuthRoutes(v1.Group("/auth"))
	routes.LinkRoutes(v1.Group("/link"))
	routes.RoleRoutes(v1.Group("/role"))

}

func main() {
	client, _ := configs.ConectToMongo()

	r := gin.Default()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.MongoInjection(client))

	// Initial Router
	initialRouter(r)

	fmt.Println("Running App on PORT")

	r.Run(":3000")
}
