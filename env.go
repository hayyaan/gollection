package gollection

import (
	"os"
	"strconv"
)

func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return fallback
}

func GetEnvInt(key string, fallback int) int {
	stringValue := GetEnv(key, "")
	if stringValue != "" {
		value, err := strconv.Atoi(stringValue)
		if err != nil {
			return fallback
		}

		return value
	}

	return fallback
}
