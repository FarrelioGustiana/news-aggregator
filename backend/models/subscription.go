package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model `json:"-"`
	ID uint `gorm:"primarykey" json:"id"`

	UserID string `gorm:"not null" json:"userId"`
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	FeedID uint `gorm:"not null" json:"feedId"`
	Feed Feed `gorm:"foreignKey:FeedID" json:"feed,omitempty"`

	SubscribedAt time.Time `json:"subscribedAt"`
}

