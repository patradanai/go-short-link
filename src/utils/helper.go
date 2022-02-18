package utils

import "net/url"

func Contains(items []string, item string) bool {
	for _, raw := range items {
		if raw == item {
			return true
		}
	}
	return false
}

func ValidateUrl(str string) bool {
	baseUrl, err := url.Parse(str)
	return err == nil && baseUrl.Host != "" && baseUrl.Scheme != ""
}
