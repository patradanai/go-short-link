package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailRequestBody struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func Authentication (c *gin.Context) {
	requestBody := EmailRequestBody{}

	// BindJSON and Validate 
	if err := c.ShouldBindJSON(&requestBody); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":"Bad Request for Authentication"})
		c.Abort()
		return 
	}

	// Fetch Username


	// Compare Password


	// Return Token


	fmt.Printf("User %v, Password %v",requestBody.Username,requestBody.Password)

}

