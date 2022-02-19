package controllers

import (
	"context"
	"net/http"
	"tiddly/src/configs"
	"tiddly/src/models"
	"tiddly/src/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type tokenRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}

func RefreshToken(c *gin.Context) {
	refreshTokenCollection := c.MustGet("mongoClient").(*mongo.Client).Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("refreshtokens")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokenFromBody := tokenRequestBody{}

	if err := c.ShouldBindJSON(&tokenFromBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	//////////////////////////////////////////////////////////
	//    				Query Single Document              	//
	//////////////////////////////////////////////////////////

	filter := bson.M{"refresh_token": tokenFromBody.RefreshToken}
	resultToken := models.RefreshToken{}
	if err := refreshTokenCollection.FindOne(ctx, filter).Decode(&resultToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "not found refresh_token"})
		return
	}

	// Check Expired Refresh Token
	if resultToken.ExpiredAt.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "expired refresh_token"})
		return
	}

	// Generation AccessToken
	token, err := utils.CreateToken(resultToken.UserId.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	data := map[string]string{
		"access_token":  token,
		"refresh_token": tokenFromBody.RefreshToken,
	}

	// Return Token
	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "authenticaion completed", "data": data})

}

func RegisterUser(c *gin.Context) {
	// Assertion Type .(*mongo.Client)
	userCollection := c.MustGet("mongoClient").(*mongo.Client).Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	requestBody := models.User{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	//////////////////////////////////////////////////////////
	//    				Insert One Document              	//
	//////////////////////////////////////////////////////////

	encryptPwd, _ := utils.EncryptBcrypt(requestBody.Password)
	userModel := models.User{Username: requestBody.Username, Password: encryptPwd, Email: requestBody.Email}
	if _, err := userCollection.InsertOne(ctx, userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "created user success"})
}

func LogoutUser(c *gin.Context) {
	refreshTokenCollection := c.MustGet("mongoClient").(*mongo.Client).Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("refreshtokens")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	refreshTokenReq := tokenRequestBody{}
	if err := c.ShouldBindJSON(&refreshTokenReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "not found refresh_token"})
		return
	}

	_, err := refreshTokenCollection.UpdateOne(ctx, bson.M{"refresh_token": refreshTokenReq.RefreshToken}, bson.M{"revoke": true})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "Logout success"})
}
