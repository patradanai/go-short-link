package main

import (
	"fmt"
	"tiddly/src/configs"
	"tiddly/src/middlewares"

	"github.com/gin-gonic/gin"
)




func main() {
	client,_ :=  configs.ConectToMongo()

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.MongoInjection(client))

	fmt.Println("Running App on PORT")

	r.Run(":3000")
}
