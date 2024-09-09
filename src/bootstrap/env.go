package bootstrap

import (
	"log"

	"github.com/joho/godotenv"
)

func initializeEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Printf("Successfully connected to environment variables")
}
