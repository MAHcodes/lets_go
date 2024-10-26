package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT   string
	DB_URL string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using default configuration")
	}

	PORT = getEnv("PORT", "8899")
	DB_URL = getEnv("DB_URL", "teendo.db")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
