package services

import (
	"errors"
	"fmt"

	"github.com/FarrelioGustiana/backend/config"
	"github.com/FarrelioGustiana/backend/models"
	"github.com/FarrelioGustiana/backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(username string, password string) (*models.User, error) {
	var existingUser models.User

	if err := config.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return nil, errors.New("username already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("database error checking existing user: %w", err)
	}
	
	hashedPasword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	user := models.User{
		Username: username,
		Password: string(hashedPasword),
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("error creating user: %w", result.Error)
	}

	return &user, nil
}

func LoginUser(username string, password string) (string, error) {
	var user models.User

	result := config.DB.Where("username = ?", username).First(&user) 
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid credentials")
		}
		return "", fmt.Errorf("database error retrieving user: %w", result.Error)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}