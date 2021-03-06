package utils

import (
	"log"

	"github.com/joho/godotenv"
)

//LoadEnvironmentVariables makes sure that we can access the environment variables in .env
func LoadEnvironmentVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
