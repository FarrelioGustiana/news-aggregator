package controllers

import (
	"net/http" // Import http package for HTTP status codes

	"github.com/FarrelioGustiana/backend/services" // Import your services package
	"github.com/gin-gonic/gin"                     // Import Gin framework
)

// AuthRequest struct defines the expected JSON payload for login and register requests.
// `binding:"required"` ensures these fields are present in the request body.
type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterUser handles the HTTP POST request for user registration.
// It binds the request JSON to AuthRequest, calls the service layer, and sends an appropriate JSON response.
func RegisterUser(c *gin.Context) {
	var req AuthRequest
	// c.ShouldBindJSON binds the request body to the struct.
	// If binding fails (e.g., missing required fields, invalid JSON), it returns an error.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400 Bad Request
		return
	}

	// Call the RegisterUser function from the services layer.
	// This separates HTTP handling from business logic.
	user, err := services.RegisterUser(req.Username, req.Password)
	if err != nil {
		// Handle specific errors from the service layer.
		if err.Error() == "user with this username already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()}) // 409 Conflict
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()}) // 500 Internal Server Error
		return
	}

	// If registration is successful, return 201 Created status with a success message and username.
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "username": user.Username})
}

// LoginUser handles the HTTP POST request for user login.
// It binds the request JSON, calls the service layer for authentication, and returns a JWT token.
func LoginUser(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400 Bad Request
		return
	}

	// Call the LoginUser function from the services layer.
	token, err := services.LoginUser(req.Username, req.Password)
	if err != nil {
		// Handle specific errors from the service layer.
		if err.Error() == "invalid credentials" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()}) // 401 Unauthorized
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login: " + err.Error()}) // 500 Internal Server Error
		return
	}

	// If login is successful, return 200 OK status with the JWT token.
	c.JSON(http.StatusOK, gin.H{"token": token})
}
