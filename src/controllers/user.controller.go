package controllers

import (
	"net/http"
	"tiddly/src/middlewares"
	"tiddly/src/models"
	"tiddly/src/repositories"
	"tiddly/src/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	userId := c.MustGet("userId").(string)

	OAuthRepository := repositories.OAuthRepository(mongoClient)

	_, exist := OAuthRepository.IsExistOAuth(userId)
	if exist {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Existing OAuth Client"})
		return
	}

	uuidNoDash := utils.GenUUIDNoDash()

	objId, _ := primitive.ObjectIDFromHex(userId)
	oAuthModel := &models.OauthClient{ApiKey: uuidNoDash, UserId: objId, Revoke: false}

	if _, err := OAuthRepository.CreateOAuth(oAuthModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "created oath client success"})

}

func GenOauthClient(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	userId := c.MustGet("userId").(string)

	OAuthRepository := repositories.OAuthRepository(mongoClient)

	uuidNoDash := utils.GenUUIDNoDash()
	objId, _ := primitive.ObjectIDFromHex(userId)

	// Inactive Last OAUTH
	criterial := bson.M{"user_id": objId, "revoke": false}
	update := bson.M{"revoke": true, "updated_at": time.Now()}
	if _, err := OAuthRepository.UpdateOAuth(criterial, update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	// Create New One

	oAuthModel := &models.OauthClient{ApiKey: uuidNoDash, UserId: objId, Revoke: false}

	if _, err := OAuthRepository.CreateOAuth(oAuthModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "updated oath client success"})
}

func RevokeOauthClient(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	userId := c.MustGet("userId").(string)
	oauthId := c.Param("oathId")

	OAuthRepository := repositories.OAuthRepository(mongoClient)

	objId, _ := primitive.ObjectIDFromHex(userId)
	objOauthId, _ := primitive.ObjectIDFromHex(oauthId)

	// Inactive Last OAUTH
	criterial := bson.M{"_id": objOauthId, "user_id": objId, "revoke": false}
	update := bson.M{"revoke": true, "updated_at": time.Now()}
	if _, err := OAuthRepository.UpdateOAuth(criterial, update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "revoked oath client success"})

}
