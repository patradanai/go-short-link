package utils

import (
	"net/url"
	"regexp"

	"github.com/google/uuid"
)

func Contains(items []string, item string) bool {
	for _, raw := range items {
		if raw == item {
			return true
		}
	}
	return false
}

func ValidateUrl(str string) error {
	_, err := url.Parse(str)
	if err != nil {
		return err
	}
	return nil
}

func GenUUIDNoDash() string {

	regex := regexp.MustCompile(`[-]`)
	uuid := uuid.New()

	uuidNoDash := regex.ReplaceAllString(uuid.String(), "")

	return uuidNoDash
}
