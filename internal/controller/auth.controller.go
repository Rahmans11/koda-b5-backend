package controller

import (
	"log"
	"net/http"

	"github.com/Rahmans11/koda-b5-backend/internal/dto"
	"github.com/Rahmans11/koda-b5-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: service.NewAuthService(),
	}
}

func (a *AuthController) PostRegister(c *gin.Context) {
	var authData dto.AuthData

	if err := c.ShouldBindJSON(&authData); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "internal server Error",
		})
		return
	}

	err := a.authService.RegisterUser(&authData)
	if err != nil {
		log.Println("Registration error:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Registration failed",
			"message": err.Error(),
		})
		return
	}

	responseData := gin.H{
		"email":   authData.Email,
		"message": "Registration successful",
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   responseData,
	})
}

func (a *AuthController) PostLogin(c *gin.Context) {
	var loginData dto.AuthData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "internal server Error",
		})
		return
	}

	err := a.authService.LoginValidation(&loginData)
	if err != nil {
		log.Println("Login error:", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Login failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login successful",
		"data": gin.H{
			"email": loginData.Email,
		},
	})
}
