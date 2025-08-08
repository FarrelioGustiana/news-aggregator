package models

import (
	"time"

	"gorm.io/gorm"
)



type Article struct {
	gorm.Model `json:"-"`
	ID          uint           `gorm:"primaryKey" json:"id"` 

	FeedID      uint           `gorm:"not null" json:"feedId"` 
	Feed        Feed           `gorm:"foreignKey:FeedID" json:"feed,omitempty"` 
	
	Title       string         `gorm:"not null;size:500" json:"title"` 
	Link        string         `gorm:"unique;not null;size:1000" json:"link"` 
	Description string         `gorm:"type:text" json:"description"` 
	PubDate     *time.Time     `json:"pubDate,omitempty"`
	GUID        string         `gorm:"unique" json:"guid"`    

	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"` 
}
