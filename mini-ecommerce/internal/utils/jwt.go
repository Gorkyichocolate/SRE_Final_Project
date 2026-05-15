package utils

import (
	"sre/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(config.LoadConfig().JWTSecret)

func GenerateToken(userID int64) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(SecretKey)
}

func ParseToken(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
}
