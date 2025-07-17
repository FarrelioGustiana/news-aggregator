package main

import (
	"log"
	"os"

	"github.com/FarrelioGustiana/backend/config"
	"github.com/FarrelioGustiana/backend/routes"
	"github.com/gin-gonic/gin"
)

// The init function is automatically executed before the main function.
// It's a good place for setup tasks like loading environment variables and connecting to the database.
func init() {
	// Load environment variables from the .env file.
	config.LoadEnv()
	// Connect to the PostgreSQL database and run GORM auto-migrations.
	config.ConnectDB()
}

// The main function is the entry point of your application.
func main() {
	// Initialize a new Gin default router.
	r := gin.Default()

	// --- Basic CORS Setup (Optional but often needed for frontend development) ---
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
	// -------------------------------------------------------------------------

	// Define a simple GET endpoint for the root path.
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pilar Credo Backend API is running!",
			"status":  "OK",
		})
	})

	// Setup API routes
	routes.SetupAPIRoutes(r)

	// TODO: The feed fetching scheduler will be started here later
	// services.StartFeedScheduler(config.DB)


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Pilar Credo Backend Server starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
