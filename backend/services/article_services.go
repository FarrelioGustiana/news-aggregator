package services

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/FarrelioGustiana/backend/config"
	"github.com/FarrelioGustiana/backend/models"
)

func GetArticlesForUser(userID string, page, pageSize int) ([]models.Article, int64, error) {
	var subscribedFeedIDs []uint
	config.DB.Model(&models.Subscription{}).
		Where("user_id = ?", userID).
		Pluck("feed_id", &subscribedFeedIDs)

	if len(subscribedFeedIDs) == 0 {
		return []models.Article{}, 0, nil
	}

	var articles []models.Article
	var totalArticles int64

	offset := (page - 1) * pageSize

	config.DB.Model(&models.Article{}).
		Where("feed_id IN (?)", subscribedFeedIDs).
		Count(&totalArticles)

	result := config.DB.Preload("Feed").
		Where("feed_id IN (?)", subscribedFeedIDs).
		Order("pub_date DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&articles)

	if result.Error != nil {
		return nil, 0, fmt.Errorf("failed to retrieve articles for user: %w", result.Error)
	}

	return articles, totalArticles, nil
}

func GetArticleByID(articleID uint, userID string) (*models.Article, error) {
	var article models.Article

	result := config.DB.Preload("Feed").
		Joins("JOIN subscriptions ON subscriptions.feed_id = articles.feed_id").
		Where("articles.id = ? AND subscriptions.user_id = ?", articleID, userID).
		First(&article)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("article not found or not subscribed")
		}
		return nil, fmt.Errorf("database error retrieving article: %w", result.Error)
	}

	return &article, nil
}
