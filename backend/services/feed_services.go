package services

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/FarrelioGustiana/backend/config"
	"github.com/FarrelioGustiana/backend/models"
)

func CreateFeed(name string, url string) (*models.Feed, error) {
	var existingFeed models.Feed
	if err := config.DB.Where("url = ?", url).First(&existingFeed).Error; err == nil {
		return nil, errors.New("feed with this URL already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("database error checking existing feed: %w", err)
	}

	feed := models.Feed{
		Name: name,
		URL: url,
	}

	result := config.DB.Create(&feed)
	if result.Error != nil {
		return nil, fmt.Errorf("error creating feed: %w", result.Error)
	}

	return &feed, nil
}

func GetAllFeeds() ([]models.Feed, error) {
	var feeds []models.Feed
	result := config.DB.Order("name ASC").Find(&feeds)
	if result.Error != nil {
		return nil, fmt.Errorf("error fetching feeds: %w", result.Error)
	}
	return feeds, nil
}