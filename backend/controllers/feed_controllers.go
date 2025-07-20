package controllers

import (
	"net/http"
	"strconv"

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

func GetFeedByID(c *gin.Context) {
	idString := c.Param("id")
	
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid feed ID format"})
		return
	}

	feed, err := services.GetFeedByID(uint(id))
	if err != nil {
		if err.Error() == "feed not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve feed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, feed)
}

// UpdateFeed handles the HTTP PUT request for updating an existing feed.
func UpdateFeed(c *gin.Context) {
	// Get the feed ID from the URL parameter.
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feed ID format"})
		return
	}

	var req FeedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feed, err := services.UpdateFeed(uint(id), req.Name, req.URL)
	if err != nil {
		if err.Error() == "feed not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) 
			return
		}
		if err.Error() == "another feed with this URL already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update feed: " + err.Error()}) 
		return
	}

	c.JSON(http.StatusOK, feed)
}

func DeleteFeed(c *gin.Context) {
	// Get the feed ID from the URL parameter.
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feed ID format"})
		return
	}

	err = services.DeleteFeed(uint(id))
	if err != nil {
		if err.Error() == "feed not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) 
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete feed: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent) 
}