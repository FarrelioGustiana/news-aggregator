package controllers

import (
	"net/http"
	"strconv"

	"github.com/FarrelioGustiana/backend/services"
	"github.com/gin-gonic/gin"
)



func GetArticlesForUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "20")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 20
	}

	articles, total, err := services.GetArticlesForUser(userID.(string), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve articles: " + err.Error()}) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func GetArticleByID(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	articleIDStr := c.Param("id")
	articleID, err := strconv.ParseUint(articleIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID format"}) // 400 Bad Request
		return
	}

	article, err := services.GetArticleByID(uint(articleID), userID.(string))
	if err != nil {
		if err.Error() == "article not found or not subscribed" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // 404 Not Found
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve article: " + err.Error()}) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, article) // 200 OK
}
