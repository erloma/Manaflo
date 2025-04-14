package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	// Parse token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	return claims.UserID, nil
}