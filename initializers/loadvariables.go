package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadVariables(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}