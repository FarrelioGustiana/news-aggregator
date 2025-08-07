package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID string `gorm:"type:uuid;unique;not null"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	CreatedAt time.Time      
	UpdatedAt time.Time      
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Subscriptions []Subscription `gorm:"foreignKeyUserID"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	return
}