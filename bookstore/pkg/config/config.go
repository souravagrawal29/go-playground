package config

import (
	"log"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variable", err)
	}
}