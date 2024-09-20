package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addCorsHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
}

func handlePreflightRequest(c *gin.Context) {
	addCorsHeaders(c)
	c.Status(http.StatusOK)
}

func main() {
	fmt.Println("Starting Gin Server")

	r := gin.Default()

	r.OPTIONS("/proxy", handlePreflightRequest)

	r.GET("/proxy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run()
}
