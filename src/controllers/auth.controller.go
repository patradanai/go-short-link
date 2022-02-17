package controllers

import (
	"context"
	"net/http"
	"strings"
	"tiddly/src/configs"
	"tiddly/src/models"
	"tiddly/src/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Authentication(c *gin.Context) {
	// Assertion Type .(*mongo.Client)
	userCollection := c.MustGet("mongoClient").(*mongo.Client).Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("users")
	refreshTokenCollection := c.MustGet("mongoClient").(*mongo.Client).Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("refreshtokens")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	requestBody := EmailRequestBody{}

	// BindJSON and Validate
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Bad Request for Authentication"})
		return
	}

	//////////////////////////////////////////////////////////
	//                                                      //
	//    				Mongo Aggregate               		//
	//                                                      //
	//////////////////////////////////////////////////////////
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "username", Value: requestBody.Username}}}}
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

	// Compare Password
	if err := utils.DecryptBcrypt(userArray[0].Password, requestBody.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "invalid user/password"})
		return
	}

	// Generation AccessToken
	token, err := utils.CreateToken(userArray[0].Id.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	getUUid := uuid.New() // Generate Refresh Token

	refreshModel := models.RefreshToken{UserId: userArray[0].Id, RefreshToken: getUUid.String(), ExpiredAt: time.Now().AddDate(0, 1, 0), CreatedAt: time.Now(), UpdatedAt: time.Now(), Revoke: false} // Insert Document
	if _, err := refreshTokenCollection.InsertOne(ctx, refreshModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	data := map[string]string{
		"access_token":  token,
		"refresh_token": getUUid.String(),
		"roles":         strings.Join(roleArray, " , "),
	}

	// Return Token
	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "authenticaion completed", "data": data})

}

type TokenRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}

func RefreshToken(c *gin.Context) {
	refreshTokenCollection := c.MustGet("mongoClient").(*mongo.Client).Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("refreshtokens")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokenFromBody := TokenRequestBody{}

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

type RegisterRequestBody struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
}

func RegisterUser(c *gin.Context) {
	// Assertion Type .(*mongo.Client)
	userCollection := c.MustGet("mongoClient").(*mongo.Client).Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	requestBody := RegisterRequestBody{}
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