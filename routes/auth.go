package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jokopriyono/basic-go-auth/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/home", controllers.Home)
	r.GET("/admin", controllers.Admin)
	r.GET("/logout", controllers.Logout)
}
