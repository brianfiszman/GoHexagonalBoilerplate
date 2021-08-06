package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
