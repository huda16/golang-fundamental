package main

import (
	"log"
	"os"
	"recipebook/database"
	"recipebook/router"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	// Get the PORT environment variable, with a fallback to a default value
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // Default port if not specified in .env
	}

	database.StartDB()
	router.StartServer().Run(":" + PORT)
}