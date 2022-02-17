package middlewares

import (
	"context"
	"net/http"
	"strings"
	"tiddly/src/configs"
	"tiddly/src/models"
	"tiddly/src/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Permission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userCollection := c.MustGet("mongoClient").(*mongo.Client).Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("users")
		userId := c.MustGet("userId").(string)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		//////////////////////////////////////////////////////////
		//    				Find Include Role              		//
		//////////////////////////////////////////////////////////
		userObjectId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Failed to parse body"})
			return
		}

		matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: userObjectId}}}}
		lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "roles"}, {Key: "localField", Value: "role_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "role_id"}}}}
		limitStage := bson.D{{Key: "$limit", Value: 1}}

		cur, err := userCollection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, limitStage}) // Exec Aggregate
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "something went wrong"})
			return
		}

		// Loop Cursor after Return
		userArray := []models.User{} // Slice with Map
		for cur.Next(ctx) {
			result := models.User{}
			if err := cur.Decode(&result); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "something went wrong"})
				return
			}
			userArray = append(userArray, result)
		}

		// Loop Role in slice
		roleArray := make([]string, 0)
		for _, role := range userArray[0].Roles {
			roleArray = append(roleArray, strings.ToUpper(role.Name))
		}

		if !utils.Contains(roleArray, permission) {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "access deniend"})
			return
		}

		c.Next()

	}
}
