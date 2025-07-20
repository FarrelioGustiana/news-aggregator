package models

import (
	"time"

	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model
	ID            uint           `gorm:"primaryKey"`
	Name          string         `gorm:"not null"`
	URL           string         `gorm:"unique;not null"`
	LastFetchedAt *time.Time     
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"` 
	
	Subscriptions []Subscription `gorm:"foreignKey:FeedID"`
}
