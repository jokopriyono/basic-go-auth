package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host         string
	Port         string
	User         string
	Password     string
	DBName       string
	SSLMode      string
	JWTSecretKey []byte
}

var DB *gorm.DB
var ENV Config

func InitDB(config Config) {
	cred := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(cred), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("Migrated database")

	DB = db
	ENV = config
}
