package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	ID uint `gorm:"primarykey"`

	UserID string `gorm:"not null"`
	User User `gorm:"foreignKey:UserID"`

	FeedID uint `gorm:"not null"`
	Feed Feed `gorm:"foreignKey:FeedID"`

	SubscribedAt time.Time
}

