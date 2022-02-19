package controllers

import (
	"net/http"
	"tiddly/src/models"
	"tiddly/src/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RoleList(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)

	roleRepository := repositories.RoleRepository(mongoClient)

	roles, err := roleRepository.ListRole()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "list of role", "data": roles})

}

func CreateRole(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	roleRepository := repositories.RoleRepository(mongoClient)

	roleRequest := models.Role{}
	if err := c.ShouldBindJSON(&roleRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	_, err := roleRepository.CreateRole(&roleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "created role success"})

}

func UpdateRole(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	roleRepository := repositories.RoleRepository(mongoClient)

	roleId := c.Param("roleId")

	roleRequest := models.Role{}
	if err := c.ShouldBindJSON(&roleRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	_, err := roleRepository.UpdateRole(roleId, &roleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "updated role success"})
}

func DeleteRole(c *gin.Context) {
	mongoClient := c.MustGet("mongoClient").(*mongo.Client)
	roleRepository := repositories.RoleRepository(mongoClient)

	roleId := c.Param("roleId")

	_, err := roleRepository.DeleteRole(roleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "deleted role success"})
}
