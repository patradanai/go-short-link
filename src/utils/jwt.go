package utils

import (
	"fmt"
	"tiddly/src/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken (userId string) (string,error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	token,err := at.SignedString([]byte(configs.LoadEnv("SECRET_KEY")))
	if err != nil {
		fmt.Println(err)
		return "",err
	}
	return token,nil
}
