package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func DbConnect() {
	// Formulate the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"))

	// Open the connection to the database using Gorm
	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Successfully connected to the database")
}


