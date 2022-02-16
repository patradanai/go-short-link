package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncryptBcrypt (raw string) (string,error) {
	toByte := []byte(raw)

	hash,err := bcrypt.GenerateFromPassword(toByte,10)
	if err != nil{
		fmt.Println(err)
	}

	return string(hash),err

}

func DecryptBcrypt (hashPwd string,plainPwd string) error {
	byteHash := []byte(hashPwd)
	bytePlain := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash,bytePlain)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}