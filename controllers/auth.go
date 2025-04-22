package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jokopriyono/basic-go-auth/models"
	"github.com/jokopriyono/basic-go-auth/utils"
)

var jwtKey = []byte("dummysecretkey")

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	models.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)
	if errHash != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	models.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
