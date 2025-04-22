package utils

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/jokopriyono/basic-go-auth/models"
)

func ParseToken(tokenString string) (claims *models.Claims, err error) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	jwtKey := []byte(models.ENV.JWTSecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
