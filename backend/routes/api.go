package routes

import (
	"net/http"

	controllers "github.com/FarrelioGustiana/backend/controllers"
	middleware "github.com/FarrelioGustiana/backend/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupAPIRoutes configures all API routes for the Gin router.
func SetupAPIRoutes(router *gin.Engine) {
	// Group for authentication-related routes (publicly accessible).
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", controllers.RegisterUser)
		authRoutes.POST("/login", controllers.LoginUser)
	}
	// Group for protected API routes (require JWT authentication).
	apiRoutes := router.Group("/api")
	// Apply the JWT authentication middleware to all routes within this group.
	apiRoutes.Use(middleware.AuthMiddleware())
	{
		// Example protected route: GET /api/protected
		// This route can only be accessed with a valid JWT token.
		apiRoutes.GET("/protected", func(c *gin.Context) {
			// Retrieve the userID from the Gin context, which was set by the AuthMiddleware.
			userID, exists := c.Get("userID")
			if !exists {
				// This case should ideally not happen if middleware works correctly, but good for safety.
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected area!", "userID": userID})
		})

		apiRoutes.POST("/feeds", controllers.CreateFeed)
		apiRoutes.GET("/feeds", controllers.GetAllFeeds)
		apiRoutes.GET("/feeds/:id", controllers.GetFeedByID)
		apiRoutes.PUT("/feeds/:id", controllers.UpdateFeed)
		apiRoutes.DELETE("/feeds/:id", controllers.DeleteFeed)
	}
}
