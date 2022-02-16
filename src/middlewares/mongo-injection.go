package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MongoInjection (db *mongo.Client, ctx context.Context) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Set("mongoClient",db)
		c.Set("mongoContxxt",ctx)

		c.Next()
	}
}