package services

import (
	"errors"
	"fmt"

	"github.com/FarrelioGustiana/backend/config"
	"github.com/FarrelioGustiana/backend/models"
	"gorm.io/gorm"
)

func SubscribeToFeed(userID string, feedID uint) (*models.Subscription, error) {
	// Make sure the user is exist
	var user models.User
	if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("database error finding user: %w", err)
	}

	// Make sure the feed is exist
	var feed models.Feed
	if err := config.DB.First(&feed, "id = ?", feedID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("feed not found")
		}
		return nil, fmt.Errorf("database error finding feed: %w", err)
	} 
	
	// Make sure the user is not subscripe to the current feed
	var existingSubsription models.Subscription
	if err := config.DB.Where("user_id = ? AND feed_id = ?", userID, feedID).First(&existingSubsription).Error; err == nil {
		return nil, errors.New("already subscribed to this feed")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("database error checking existing subscription: %w", err)
	}

	subcription := models.Subscription{
		UserID: userID,
		FeedID: feedID,
		User: user,
		Feed: feed,
	}
	
	result := config.DB.Create(&subcription)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create subsription: %w", result.Error)
	}

	return &subcription, nil
}

func GetUserSubscriptions(userID string) ([]models.Subscription, error) {
	var subcriptions []models.Subscription

	result := config.DB.Preload("Feed").Where("user_id = ?", userID).Find(&subcriptions)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve subscriptions: %w", result.Error)
	}

	return subcriptions, nil
}

func UnsubscribeFromFeed(userID string, feedID uint) error {
	result := config.DB.Where("user_id = ? AND feed_id = ?", userID, feedID).Delete(&models.Subscription{})

	if result.Error != nil {
		return fmt.Errorf("failed to delete subscription: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("subscription not found")
	}
	
	return nil
}

func IsUserSubscribed(userID string, feedID uint) (bool, error) {
	var count int64
	err := config.DB.Model(&models.Subscription{}).
		Where("user_id = ? AND feed_id = ?", userID, feedID).
		Count(&count).Error

	if err != nil {
		return false, fmt.Errorf("database error checking subscription status: %w", err)
	}

	return count > 0, nil
}