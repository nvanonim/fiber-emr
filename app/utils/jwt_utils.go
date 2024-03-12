package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTExpirationTime is the expiration time in minutes
var JWTExpirationTime = 24 * time.Hour

// CreateToken creates a jwt token
func CreateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iss": "fiber-emr",
		"exp": time.Now().Add(JWTExpirationTime).Unix(),
	})

	secretKey := GetEnv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetJWTClaims returns the claims of the jwt token
func GetJWTClaims(c *gin.Context) jwt.MapClaims {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetEnv("JWT_SECRET")), nil
	})

	return token.Claims.(jwt.MapClaims)
}
