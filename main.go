package main

import (
	"code-engine/database"
	"code-engine/models"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.Language{})
	database.Database.AutoMigrate(&models.Submission{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
