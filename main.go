package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jokopriyono/basic-go-auth/models"
	"github.com/jokopriyono/basic-go-auth/routes"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.Config{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		SSLMode:      os.Getenv("DB_SSLMODE"),
		JWTSecretKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}
	models.InitDB(config)
	routes.AuthRoutes(r)
	r.Run(":8080")
}
