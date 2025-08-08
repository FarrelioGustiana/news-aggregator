package models

import (
	"time"

	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model `json:"-"`
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"not null" json:"name"`
	URL           string         `gorm:"unique;not null" json:"url"`
	LastFetchedAt *time.Time     `json:"lastFetchedAt,omitempty"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"` 
	
	Subscriptions []Subscription `gorm:"foreignKey:FeedID" json:"subscriptions,omitempty"`
}
