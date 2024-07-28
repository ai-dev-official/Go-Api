package main

import (
	"golang.org/aidev/crud-in-go/initializers"
	"golang.org/aidev/crud-in-go/models"
)


func init() {
	initializers.LoadVariables()
	initializers.DbConnect()
}

func main() {
	// Migrate the schema
  initializers.DB.AutoMigrate(&models.Product{})
}