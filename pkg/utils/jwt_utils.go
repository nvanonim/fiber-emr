package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iss": "fiber-emr",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := GetEnv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
