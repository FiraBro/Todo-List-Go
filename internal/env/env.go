package env

import (
	"os"
	"strconv"
)

// GetString reads a string environment variable,
// returns defaultValue if not found.
func GetString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetInt reads an integer environment variable,
// returns defaultValue if not found or invalid.
func GetInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.Atoi(value)
		if err == nil {
			return i
		}
	}
	return defaultValue
}
