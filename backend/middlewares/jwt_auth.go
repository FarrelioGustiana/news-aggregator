package middleware

import (
	"net/http" // Import http package for HTTP status codes
	"strings"  // Import strings package for string manipulation

	"github.com/FarrelioGustiana/backend/utils" // Import your utils package for JWT functions
	"github.com/gin-gonic/gin"                  // Import Gin framework
)

// AuthMiddleware is a Gin middleware function that authenticates requests using JWT.
// It checks for a valid JWT in the "Authorization" header and sets the user ID in the Gin context.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header from the request.
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// If no Authorization header is present, return 401 Unauthorized.
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort() // Abort the request chain
			return
		}

		// The Authorization header typically has the format "Bearer TOKEN_STRING".
		// Split the header value to extract the token string.
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			// If the format is invalid, return 401 Unauthorized.
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format. Expected 'Bearer TOKEN'"})
			c.Abort()
			return
		}

		// Extract the actual token string.
		tokenString := parts[1]

		// Validate the token using the utility function.
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			// If token validation fails (e.g., invalid signature, expired token), return 401 Unauthorized.
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token: " + err.Error()})
			c.Abort()
			return
		}

		// Extract the user ID from the token claims.
		// JWT claims numeric values often come as float64, so we cast it to string (UUID).
		userID, ok := (*claims)["user_id"].(string) // User ID is a string (UUID)
		if !ok {
			// If user_id claim is missing or not a string, return 401 Unauthorized.
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload: user ID not found or invalid type"})
			c.Abort()
			return
		}

		// Set the user ID in the Gin context. This makes the user ID accessible
		// to subsequent handlers in the request chain (e.g., controllers).
		c.Set("userID", userID)

		// Proceed to the next middleware or the actual route handler.
		c.Next()
	}
}
