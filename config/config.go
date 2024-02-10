package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	// Load environment variables from .env file if present
	err := godotenv.Load(".env")
	if err != nil {
		// If .env file not found, load from OS environment
		log.Println(".env file not found, loading from OS environment")
	}
}
