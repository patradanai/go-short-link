package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"tiddly/src/configs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthHeader struct {
	Token string `header:"Authorization"`
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerAuth := AuthHeader{}

		if err := c.BindHeader(&headerAuth); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorize"})
			return
		}

		getToken := strings.Split(headerAuth.Token, "Bearer ")

		if len(getToken) < 2 {
			c.JSON(http.StatusUnauthorized,
				"Must provide Authorization header with format `Bearer {token}")
			return
		}

		// Vertify Token
		token, err := jwt.Parse(getToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}

			return []byte(configs.LoadEnv("SECRET_KEY")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity,
				gin.H{"success": false, "message": err.Error()})
			return
		}

		if token.Valid {
			// AssertType jwt.MapClaims
			claims := token.Claims.(jwt.MapClaims)

			c.Set("userId", claims["userId"])
			c.Next()
		}

	}
}
