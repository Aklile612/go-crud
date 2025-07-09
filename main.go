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

	r.POST("/post", controllers.PostCreate)
	r.GET("/getpost", controllers.PostIndex)
	r.GET("/singlepost/:id", controllers.PostShow)
	r.PUT("/updatepost/:id", controllers.PostUpdate)
	r.DELETE("/deletepost/:id", controllers.PostDelete)

	r.Run() 
}
