package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GET_ENV(key string, defaultValue string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  if value, ok := os.LookupEnv(key); ok {
    return value
  }

  return defaultValue
}