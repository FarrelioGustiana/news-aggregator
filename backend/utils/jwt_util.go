package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string, isAdmin bool) (string, error) {
	// Define the claims (payload) for the JWT.
	// "authorized": A custom claim indicating if the user is authorized.
	// "user_id": The ID of the user, stored as a string.
	// "is_admin": Boolean indicating if the user has admin privileges.
	// "exp": Expiration time (Unix timestamp). Token expires in one month.
	// "iat": Issued at time (Unix timestamp).
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"is_admin":   isAdmin,
		"exp":        time.Now().Add(time.Hour * 24 * 30).Unix(), // Token expires in one month
		"iat":        time.Now().Unix(),
	}

	// Create a new JWT token with the HS256 signing method and the defined claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key from environment variables.
	// The secret key must be kept confidential and should be strong.
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err // Return error if signing fails
	}
	return signedToken, nil // Return the signed token
}

func ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid 
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid and extract the claims.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil // Return the claims if token is valid
	}

	// If the token is not valid or claims cannot be extracted, return a generic error.
	return nil, jwt.ErrInvalidKey
}
