package middleware

import (
	"net/http"
	"strings"

	"github.com/FarrelioGustiana/backend/utils"
	"github.com/gin-gonic/gin"
)

// AdminMiddleware is a Gin middleware function that verifies a user has admin privileges.
// It should be used after the regular AuthMiddleware to ensure the user is authenticated first.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header from the request.
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Split the header to extract the token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// Extract the token string
		tokenString := parts[1]

		// Validate the token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Check if the user is an admin
		isAdmin, ok := (*claims)["is_admin"].(bool)
		if !ok || !isAdmin {
			// If not admin or the claim is missing, return 403 Forbidden
			c.JSON(http.StatusForbidden, gin.H{"error": "Administrator privileges required for this action"})
			c.Abort()
			return
		}

		// User is authenticated and has admin privileges, proceed
		c.Next()
	}
}
