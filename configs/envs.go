package configs

import (
	"log"
	"os"
	"strconv"
)

func getStringEnv(key string) string {
	result := os.Getenv(key)

	if result == "" {
		log.Fatalf("%s can't be empty", key)
	}

	return result
}

func getIntEnv(key string) int {
	result, err := strconv.Atoi(os.Getenv(key))

	if err != nil {
		log.Fatalf("%s must be int", key)
	}

	return result
}
