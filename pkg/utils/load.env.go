package utils

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envPath := filepath.Join("..", ".env")

	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
