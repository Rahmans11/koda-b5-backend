package router

import (
	"github.com/Rahmans11/koda-b5-backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	authController := controller.NewAuthController()

	app.POST("/auth/register", authController.PostRegister)
	app.POST("/auth", authController.PostLogin)
}
