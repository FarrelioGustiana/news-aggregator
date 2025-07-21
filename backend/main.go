package main

import (
	"log"
	"os"

	"github.com/FarrelioGustiana/backend/config"
	"github.com/FarrelioGustiana/backend/routes"
	"github.com/FarrelioGustiana/backend/services"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pilar Credo Backend API is running!",
			"status":  "OK",
		})
	})

	routes.SetupAPIRoutes(r)

	services.StartFeedScheduler(config.DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Pilar Credo Backend Server starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
