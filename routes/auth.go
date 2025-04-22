package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jokopriyono/basic-go-auth/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.SignUp)
}
