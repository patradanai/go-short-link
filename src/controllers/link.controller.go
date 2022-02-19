package controllers

import (
	"net/http"
	"tiddly/src/configs"
	"tiddly/src/models"
	"tiddly/src/repositories"
	"tiddly/src/utils"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type createLinkRequest struct {
	Destination string `json:"destination"`
	Title       string `json:"title"`
	SlagTag     string `json:"slag_tag"`
}

func CreateLink(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	linkRepository := repositories.LinkRepository(mongoClient)

	shortLinkReq := createLinkRequest{}

	if err := c.ShouldBindJSON(&shortLinkReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Validate URl
	if err := utils.ValidateUrl(shortLinkReq.Destination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": true, "message": "url is invalid format"})
		return
	}

	// Create Short link
	uniqueCode, err := gonanoid.New(9)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}
	// Check Long Url
	filter := bson.M{"original_url": shortLinkReq.Destination}
	if getFindLink, err := linkRepository.GetLink(filter); err == nil {
		data := map[string]string{
			"title":       getFindLink.Title,
			"shortUrl":    getFindLink.ShortUrl,
			"originalUrl": getFindLink.OriginalUrl,
			"expiredAt":   getFindLink.ExpiredAt.String(),
		}

		// Response
		c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "exising original url", "data": data})
		return
	}
	// Insert to db
	shortUrl := configs.LoadEnv("BASE_URL") + "/" + uniqueCode
	shortLinkModel := models.ShortLink{Title: shortLinkReq.Title, OriginalUrl: shortLinkReq.Destination, RefCode: uniqueCode, ShortUrl: shortUrl, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	_, err = linkRepository.CreateLink(shortLinkModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}
	data := map[string]string{
		"title":       shortLinkModel.Title,
		"shortUrl":    shortLinkModel.ShortUrl,
		"originalUrl": shortLinkModel.OriginalUrl,
		"expiredAt":   shortLinkModel.ExpiredAt.String(),
	}

	// Response
	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "created completed short url", "data": data})
}

func RedirectLink(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)

	linkRepository := repositories.LinkRepository(mongoClient)

	shortId := c.Param("shortId")

	filter := bson.M{"ref_code": shortId}
	result, err := linkRepository.GetLink(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": true, "message": err.Error()})
		return
	}

	// Response Redirect
	c.Redirect(302, result.OriginalUrl)
}
