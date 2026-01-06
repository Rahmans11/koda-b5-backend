package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	var authData AuthData
	app.POST("/register", func(c *gin.Context) {

		if err := c.ShouldBindJSON(&authData); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "internal server Error",
			})
			return
		}

		if !strings.Contains(authData.Email, "@") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
			return
		}
		if strings.HasPrefix(authData.Email, "@") || strings.HasSuffix(authData.Email, "@") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
			return
		}
		parts := strings.Split(authData.Email, "@")
		if len(parts) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
			return
		}
		domain := parts[1]
		if !strings.Contains(domain, ".") || strings.HasPrefix(domain, ".") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
			return
		}

		var hasUpper, hasLower, hasDigit, hasSpecial bool

		if len(authData.Password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "minimum password 6 character"})
			return
		}

		for _, char := range authData.Password {
			if unicode.IsUpper(char) {
				hasUpper = true
			} else if unicode.IsLower(char) {
				hasLower = true
			} else if unicode.IsDigit(char) {
				hasDigit = true
			} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
				hasSpecial = true
			}
		}

		if !hasUpper {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password must have one uppercase letter"})
			return
		}
		if !hasLower {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password must have one lowercase letter"})
			return
		}
		if !hasDigit {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password must have one digit number"})
			return
		}
		if !hasSpecial {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password must have one special character"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":  "Success to Register",
			"data": authData,
		})
		fmt.Println(authData)
	})

	var loginData AuthData
	app.POST("/auth", func(c *gin.Context) {

		if err := c.ShouldBindJSON(&loginData); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "internal server Error",
			})
			return
		}
		if loginData != authData {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Success to login",
			"data": "Wellcome " + loginData.Email,
		})
	})
	app.Run("localhost:8080")
}

type AuthData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
