package services

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("freshsalmon")

func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to extract claims from token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username not found in token claims")
	}

	return username, nil
}

func ProtectedHandler(c *gin.Context) (string, error) {
	c.Header("Content-Type", "application/json")
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return "", fmt.Errorf("Missing authorization header")
	}
	tokenString = tokenString[len("Bearer "):]

	username, err := verifyToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("Invalid token")
	}
	return username, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
