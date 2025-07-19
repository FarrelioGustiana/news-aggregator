package controllers

import (
	"net/http"

	"github.com/FarrelioGustiana/backend/services"
	"github.com/gin-gonic/gin"
)

type FeedRequest struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"url" binding:"required,url"`
}

func CreateFeed(c *gin.Context) {
	var request FeedRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{"error": err.Error()},
		)
		// 400 Bad Request
		return
	}

	feed, err := services.CreateFeed(request.Name, request.URL)
	if err != nil {
		if err.Error() == "feed with this URL already exists" {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create feed: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, feed)
}

func GetAllFeeds(c *gin.Context) {
	feeds, err := services.GetAllFeeds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve feeds: " + err.Error()})
	}

	c.JSON(http.StatusOK, feeds)
}