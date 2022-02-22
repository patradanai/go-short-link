package controllers

import (
	"net/http"
	"tiddly/src/middlewares"
	"tiddly/src/models"
	"tiddly/src/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type File struct {
	File middlewares.FileUpload
}

func UpdateUserInfo(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	userId := c.MustGet("userId").(string)

	userRepository := repositories.UserRepository(mongoClient)

	// Fetch User info
	userInfo, err := userRepository.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	userInfoRequest := models.UserInfo{}
	if err := c.ShouldBind(&userInfoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Bind With File from Middleware
	if value, exist := c.Get("file"); exist {
		file := value.(*middlewares.FileUpload)
		userInfoRequest.Image.Src = file.Path
		userInfoRequest.Image.Title = file.Filename
	} else {
		userInfoRequest.Image.Src = userInfo.UserInfo.Image.Src
		userInfoRequest.Image.Title = userInfo.UserInfo.Image.Title
	}

	userInfoNew := bson.M{"$set": bson.M{"user_info": userInfoRequest}}
	_, err = userRepository.UpdateUserInfo(userId, userInfoNew)
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
