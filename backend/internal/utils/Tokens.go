package utils

import (
	"fmt"
	modelPG "lyked-backend/internal/models/postgresql"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Tokens.go
func GenerateToken(userID string, db *gorm.DB) (string, error) {
	var user modelPG.User
	var token string
	if userID == "" {
		return "", fmt.Errorf("UserID cannot be empty")
	}

	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	if user.ID.String() == "" {
		return "", fmt.Errorf("userID not found in database")
	}
	// Use bcrypt to hash the userID as a token

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.ID.String()), 14)
	if err != nil {

		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	token = string(bytes)
	return token, nil
}

func ValidateToken(token string) (string, error) {
	return "", nil // Placeholder for token validation logic
}

func RefreshToken(token string) (string, error) {
	return "", nil // Placeholder for token refresh logic
}
