package models

import (
	"time"

	"gorm.io/gorm"
)



type Article struct {
	gorm.Model
	ID          uint           `gorm:"primaryKey"` 

	FeedID      uint           `gorm:"not null"` 
	Feed        Feed           `gorm:"foreignKey:FeedID"` 
	
	Title       string         `gorm:"not null;size:500"` 
	Link        string         `gorm:"unique;not null;size:1000"` 
	Description string         `gorm:"type:text"` 
	PubDate     *time.Time     
	GUID        string         `gorm:"unique"`    

	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"` 
}
