package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID string `gorm:"type:uuid;unique;not null" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Subscriptions []Subscription `gorm:"foreignKeyUserID" json:"subscriptions,omitempty"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	return
}