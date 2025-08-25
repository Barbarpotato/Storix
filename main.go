package main

import (
	"net/http"

	"github.com/Barbarpotato/Storix/configs"
	"github.com/gin-gonic/gin"
)

func main() {

	configs.ConnectDB()

	r := gin.Default()

	// your existing endpoints remain
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Gin 🚀")
	})

	r.GET("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Hello from Gin",
		})
	})

	r.Run(":8080")
}
