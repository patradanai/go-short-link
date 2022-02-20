package controllers

import (
	"net/http"
	"tiddly/src/models"
	"tiddly/src/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type fileUpload struct {
	original string
	filename string
	path     string
}

func UpdateUserInfo(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	// file := c.MustGet("file").(*fileUpload)
	userId := c.MustGet("userId").(string)

	userRepository := repositories.UserRepository(mongoClient)

	userInfoRequest := models.UserInfo{}
	if err := c.ShouldBindJSON(&userInfoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	userInfo := bson.M{"$set": bson.M{"user_info": userInfoRequest}}
	_, err := userRepository.UpdateUserInfo(userId, userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong during update user"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "already updated user"})

}

func CreateOauthClient(c *gin.Context) {

}

func UpdateOauthClient(c *gin.Context) {

}

func RevokeOauthClient(c *gin.Context) {}
