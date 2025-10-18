package utils

import (
	"os"
)

func GET_ENV(key string, defaultValue string) string {

  // // load .env file
  // err := godotenv.Load(".env")

  // if err != nil {
  //   log.Fatalf("Error loading .env file")
  // }

  // if value, ok := os.LookupEnv(key); ok {
  //   return value
  // }

  value := os.Getenv(key)

  if value != "" {
    return value
  }

  return defaultValue
}