package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	var strongPasswordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	var authData AuthData
	app.POST("/auth", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&authData); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "internal server Error",
			})
			return
		}
		email := authData.email
		password := authData.password
		if emailRegex.MatchString(email) {
			if strongPasswordRegex.MatchString(password) {
				c.JSON(http.StatusOK, gin.H{
					"msg":  "Success to Register",
					"data": authData,
				})
			}
		}
	})

	app.Run("localhost:8080")
}

type AuthData struct {
	email    string
	password string
}
