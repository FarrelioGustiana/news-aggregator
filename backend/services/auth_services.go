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
		return nil, errors.New("user with this username already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("database error checking existing user: %w", err)
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		fmt.Printf("DEBUG: Error creating user: %v\n", result.Error)
		return nil, fmt.Errorf("error creating user: %w", result.Error)
	}

	fmt.Printf("DEBUG: User created successfully - ID: %v, Username: %s\n", user.ID, user.Username)
	return &user, nil
}

func LoginUser(username string, password string) (string, error) {
	var user models.User

	result := config.DB.Where("username = ?", username).First(&user) 
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Printf("DEBUG: User not found during login - Username: %s\n", username)
			return "", errors.New("invalid credentials")
		}
		fmt.Printf("DEBUG: Database error during login: %v\n", result.Error)
		return "", fmt.Errorf("database error retrieving user: %w", result.Error)
	}

	fmt.Printf("DEBUG: User found during login - ID: %v, Username: %s\n", user.ID, user.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}