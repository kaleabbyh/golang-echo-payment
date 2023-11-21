package utils

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var secretKey string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	secretKey     = os.Getenv("SECRETKEY")
}


func GenerateToken(userID uuid.UUID,role string) (string, error) {
	secretKey := []byte(secretKey)
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func ValidateToken(tokenString string) (uuid.UUID,string, error) {
	secretKey := []byte(secretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return uuid.UUID{},"", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{},"", errors.New("invalid token claims")
	}



	userIDStr, ok := claims["user_id"].(string)
	role,ok := claims["role"].(string)
	if !ok {
		return uuid.UUID{},"", errors.New("invalid user ID or role in token claims")
	}
	
	
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.UUID{},"", errors.New("failed to parse user ID from token claims")
	}

	return userID,role, nil
}



func GetValuesFromToken(c echo.Context) (uuid.UUID, string, error) {
	token := c.Request().Header.Get("Authorization")

	if token == "" {
		token = c.QueryParam("token")
	}

	tokenParts := strings.Split(token, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return uuid.UUID{}, "", nil
	}

	token = tokenParts[1]
	userID, role, _ := ValidateToken(token)
	return userID, role, nil
}


