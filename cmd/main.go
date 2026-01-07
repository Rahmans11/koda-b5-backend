package main

import (
	"github.com/Rahmans11/koda-b5-backend/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	router.Init(app)

	app.Run("localhost:8080")
}
