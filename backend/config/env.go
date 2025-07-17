package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return err
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("Environment variable " + key + " is not set")
	}
	return value
}