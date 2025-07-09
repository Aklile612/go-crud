package main

import (
	"github.com/Aklile612/go-crud/initalizers"
	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectionToDB()
}
func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // Default runs on :8080
}
