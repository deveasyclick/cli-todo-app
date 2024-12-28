package loaders

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
