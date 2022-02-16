package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MongoInjection (db *mongo.Client) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Set("mongoClient",db)
	
		c.Next()
	}
}