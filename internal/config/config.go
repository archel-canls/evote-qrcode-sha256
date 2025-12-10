package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads .env file once at startup
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env not found, using system environment variables")
	}
}

// GetEnv helper
func GetEnv(key string) string {
	return os.Getenv(key)
}
