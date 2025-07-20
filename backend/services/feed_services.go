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

func GetFeedByID(id uint) (*models.Feed, error) {
	var feed models.Feed

	result := config.DB.First(&feed, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("feed not found")
		}

		return nil, fmt.Errorf("database error retrieving feed: %w", result.Error)
	}

	return &feed, nil
}

func UpdateFeed(id uint, newName string, newURL string) (*models.Feed, error) {
	var feed models.Feed

	result := config.DB.First(&feed, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("feed not found")
		}

		return nil, fmt.Errorf("database error retrieving feed: %w", result.Error)
	}

	// Check for duplicate URL (other than this feed)
	var existingFeedURL models.Feed
	if err := config.DB.Where("id <> ? AND url = ?", id, newURL).First(&existingFeedURL).Error; err == nil {
		// If we found a record, it means a duplicate exists
		return nil, errors.New("another feed with this URL already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// If error is not 'record not found', it's a database error
		return nil, fmt.Errorf("database error checking for duplicate URL: %w", err)
	}

	feed.Name = newName
	feed.URL = newURL

	result = config.DB.Save(&feed)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update feed: %w", result.Error)
	}

	return &feed, nil
}

func DeleteFeed(id uint) error {
	result := config.DB.Delete(&models.Feed{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete feed: %w", result.Error)
	} 

	if result.RowsAffected == 0 {
		return errors.New("feed not found")
	}

	return nil
}