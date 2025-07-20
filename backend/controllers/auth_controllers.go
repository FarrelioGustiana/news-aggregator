package controllers

import (
	"net/http"

	"github.com/FarrelioGustiana/backend/services"
	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400 Bad Request
		return
	}

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

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "username": user.Username})
}

func LoginUser(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400 Bad Request
		return
	}

	token, err := services.LoginUser(req.Username, req.Password)
	if err != nil {
		if err.Error() == "invalid credentials" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()}) // 401 Unauthorized
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login: " + err.Error()}) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
