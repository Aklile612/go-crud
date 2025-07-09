package controllers

import (
	"github.com/Aklile612/go-crud/initalizers"
	"github.com/Aklile612/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	//get data from request body

	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)
	//create a post

	post:=models.Post{Title: body.Title, Body: body.Body}
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


func PostIndex(c *gin.Context){
	var posts []models.Post
	initalizers.DB.Find(&posts)


	c.JSON(200,gin.H{
		"Posts":posts,
	})
}
