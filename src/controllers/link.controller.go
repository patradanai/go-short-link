package controllers

import (
	"context"
	"fmt"
	"net/http"
	"tiddly/src/configs"
	"tiddly/src/models"
	"tiddly/src/utils"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateLinkRequest struct {
	Destination string `json:"destination"`
	Title       string `json:"title"`
	SlagTag     string `json:"slag_tag"`
}

func CreateLink(c *gin.Context) {
	shortLinkCollection := c.MustGet("mongoClient").(*mongo.Client).Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("shortlinks")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	shortLinkReq := CreateLinkRequest{}

	if err := c.ShouldBindJSON(&shortLinkReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Validate URl
	if !utils.ValidateUrl(shortLinkReq.Destination) {
		c.JSON(http.StatusBadRequest, gin.H{"success": true, "message": "url is invalid format"})
		return
	}

	// Create Short link
	uniqueCode, err := gonanoid.New()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}
	// Check Long Url
	filter := bson.M{"original_url": shortLinkReq.Destination}
	result := models.ShortLink{}
	if err := shortLinkCollection.FindOne(ctx, filter).Decode(&result); err == nil {
		data := map[string]string{
			"title":       "",
			"shortUrl":    result.ShortUrl,
			"originalUrl": result.OriginalUrl,
			"expiredAt":   "",
		}

		// Response
		c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "exising original url", "data": data})
		return
	}
	// Insert to db
	shortUrl := configs.LoadEnv("BASE_URL") + "/" + uniqueCode
	shortLinkModel := models.ShortLink{OriginalUrl: shortLinkReq.Destination, RefCode: uniqueCode, ShortUrl: shortUrl, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if _, err := shortLinkCollection.InsertOne(ctx, shortLinkModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}
	data := map[string]string{
		"title":       "",
		"shortUrl":    shortUrl,
		"originalUrl": shortLinkReq.Destination,
		"expiredAt":   "",
	}

	// Response
	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "created completed short url", "data": data})
}

type Params struct {
	ShortRefId string `uri:"shortId"`
}

func RedirectLink(c *gin.Context) {
	shortLinkCollection := c.MustGet("mongoClient").(*mongo.Client).Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("shortlinks")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	shortId := Params{}

	if err := c.BindUri(&shortId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": true, "message": err.Error()})
		return
	}

	fmt.Println(shortId.ShortRefId)

	filter := bson.M{"ref_code": shortId.ShortRefId}
	result := models.ShortLink{}
	if err := shortLinkCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": true, "message": err.Error()})
		return
	}

	// Response Redirect
	c.Redirect(302, result.OriginalUrl)
}
