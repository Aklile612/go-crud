package controllers

import (
	"github.com/Aklile612/go-crud/initalizers"
	"github.com/Aklile612/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	//get data from request body

	//create a post

	post:=models.Post{Title: "hello", Body: "Post the body"}
	result:= initalizers.DB.Create(&post)


	if result.Error != nil{
		c.Status(400)
		return
	}
	//return it
	c.JSON(200, gin.H{
		"Post": post,
	})
}
