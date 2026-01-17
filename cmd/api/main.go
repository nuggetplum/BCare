package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nuggetplum/BCare/internal/database"
)

func main() {
	// 1. Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Connect to Database
	database.ConnectDB()

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
