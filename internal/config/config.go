package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DatabaseUrl        string
	TokenEncryptionKey string
	AuthFileName       string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DatabaseUrl = os.Getenv("DATABASE_URL")
	TokenEncryptionKey = os.Getenv("TOKEN_ENCRYPTION_KEY")
	home, _ := os.UserHomeDir()
	AuthFileName = home + os.Getenv("AUTH_FILENAME")
}
