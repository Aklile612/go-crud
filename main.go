package main

import (
	"github.com/Aklile612/go-crud/controllers"
	"github.com/Aklile612/go-crud/initalizers"
	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectionToDB()
}
func main() {
	r := gin.Default()

	r.GET("/", controllers.PostCreate)

	r.Run() // Default runs on :8080
}
