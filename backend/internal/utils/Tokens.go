package utils

import (
	"errors"
	"fmt"
	jwtModel "lyked-backend/internal/models/jwt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(GetEnv("JWT_SECRET_KEY", "")) // Replace with your secret key

// Tokens.go
func GenerateToken(userID string, email string, username string) (string, error) {
	experationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
	claims := jwtModel.JWTClaims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(experationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "lyked-app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ValidateToken(token string) (*jwtModel.JWTClaims, error) {
	claims := &jwtModel.JWTClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func RefreshToken(token string) (string, error) {
	claims, err := ValidateToken(token)
	if err != nil {
		return "", fmt.Errorf("failed to validate token: %w", err)
	}

	if time.Until(claims.ExpiresAt.Time) > 1*time.Hour {
		return "", fmt.Errorf("token not eligible for refresh yet")
	}
	return GenerateToken(claims.ID, claims.Email, claims.Username)
}
