package utils

func Contains(items []string, item string) bool {
	for _, raw := range items {
		if raw == item {
			return true
		}
	}
	return false
}
