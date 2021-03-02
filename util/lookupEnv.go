package util

import (
	"os"
)

// looks up an environment variable and returns the var or a fallback
func LookupEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LookupEnvBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if value == "true" {
			return true
		}
		return false
	}
	return fallback
}
