package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nuggetplum/BCare/internal/database"
	"github.com/nuggetplum/BCare/internal/models"
)

func main() {
	// 1. Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Connect to Database
	database.ConnectDB()

	database.DB.AutoMigrate(&models.Product{})

	// 3. Setup Router
	r := gin.Default()

	// 4. Health Check Endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  "database connected",
		})
	})

	// 5. Run Server
	log.Println("Server running on port 8080")
	r.Run(":8080")
}

func seedProducts() {
	var count int64
	database.DB.Model(&models.Product{}).Count(&count)

	if count == 0 {
		products := []models.Product{
			{Name: "Go Programming Blueprints", Description: "Learn Go by building real apps", Price: 39.99, ImageURL: "https://via.placeholder.com/150", Stock: 10},
			{Name: "Mechanical Keyboard", Description: "Clicky switches for coding", Price: 120.00, ImageURL: "https://via.placeholder.com/150", Stock: 5},
			{Name: "Noise Cancelling Headphones", Description: "Focus mode enabled", Price: 299.50, ImageURL: "https://via.placeholder.com/150", Stock: 8},
		}
		database.DB.Create(&products)
		log.Println("🌱 Seeded dummy products")
	}
}
