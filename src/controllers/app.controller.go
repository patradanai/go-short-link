package controllers

import (
	"net/http"
	"tiddly/src/models"
	"tiddly/src/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ListAppPackage(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)

	packageRepository := repositories.AppRepository(mongoClient)

	// Get Packages
	packages, err := packageRepository.ListAppPackage()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "list of packages", "data": packages})
}

func GetAppPackage(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	packageId := c.Param("packageId")

	packageRepository := repositories.AppRepository(mongoClient)

	// Get Package
	result, err := packageRepository.GetAppPackage(packageId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "package", "data": result})

}

func CreateAppPackage(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)

	packageRepository := repositories.AppRepository(mongoClient)

	packageRequest := models.AppPackage{}

	// Bind Request Body
	if err := c.ShouldBindJSON(&packageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	// Created Package
	if _, err := packageRepository.CreateAppPackage(&packageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "package", "data": "created package success"})

}

func DeleteAppPackage(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	packageId := c.Param("packageId")

	packageRepository := repositories.AppRepository(mongoClient)

	if _, err := packageRepository.DeleteAppPackage(packageId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "package", "data": "deleted package success"})

}
