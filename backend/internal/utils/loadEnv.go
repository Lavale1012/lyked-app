package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Println("No .env file found — using system environment variables")
	} else {
		log.Println("✅ Environment variables loaded from .env file")
	}
}

func GetEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
