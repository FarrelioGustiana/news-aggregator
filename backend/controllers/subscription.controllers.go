package controllers

import (
	"net/http"
	"strconv"

	"github.com/FarrelioGustiana/backend/services"
	"github.com/gin-gonic/gin"
)

type SubscribeRequest struct {
	FeedID uint `json:"feed_id" binding:"required"`
}

type SubscriptionResponse struct {
	SubscriptionID uint   `json:"subscription_id"`
	FeedID         uint   `json:"feed_id"`
	FeedName       string `json:"feed_name"`
	FeedURL        string `json:"feed_url"`
}

func SubscribeToFeed(c *gin.Context) {
	// Get the user id from the Auth Middleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	var req SubscribeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400 Bad Request
		return
	}

	subscription, err := services.SubscribeToFeed(userID.(string), uint(req.FeedID)) 
	if err != nil {
		if err.Error() == "user not found" || err.Error() == "feed not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // 404 Not Found
			return
		}
		if err.Error() == "already subscribed to this feed" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()}) // 409 Conflict
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to subscribe: " + err.Error()}) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusCreated, subscription) // 201 Created
}

func GetUserSubscriptions(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	subscriptions, err := services.GetUserSubscriptions(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve subscriptions: " + err.Error()}) // 500 Internal Server Error
		return
	}

	var response []SubscriptionResponse
	for _, sub := range subscriptions {
		response = append(response, SubscriptionResponse{
			SubscriptionID: sub.ID,
			FeedID:         sub.Feed.ID,
			FeedName:       sub.Feed.Name,
			FeedURL:        sub.Feed.URL,
		})
	}

	c.JSON(http.StatusOK, response) // 200 OK
}

func UnsubscribeFromFeed(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	feedIDStr := c.Param("feedId")
	feedID, err := strconv.ParseUint(feedIDStr, 10, 32) // Parse as uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feed ID format"}) // 400 Bad Request
		return
	}

	err = services.UnsubscribeFromFeed(userID.(string), uint(feedID))
	if err != nil {
		if err.Error() == "subscription not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // 404 Not Found
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unsubscribe: " + err.Error()}) // 500 Internal Server Error
		return
	}

	c.Status(http.StatusNoContent) // 204 No Content for successful deletion
}

func CheckSubscriptionStatus(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	feedIDStr := c.Param("feedId")
	feedID, err := strconv.ParseUint(feedIDStr, 10, 32) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feed ID format"}) // 400 Bad Request
		return
	}

	isSubscribed, err := services.IsUserSubscribed(userID.(string), uint(feedID)) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check subscription status: " + err.Error()}) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, gin.H{"is_subscribed": isSubscribed}) // 200 OK
}
