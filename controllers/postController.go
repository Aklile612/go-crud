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
func PostShow(c *gin.Context){
	//get id off url
	id:= c.Param("id")

	//get the post
	var post models.Post
	initalizers.DB.First(&post,id)


	c.JSON(200,gin.H{
		"Posts":post,
	})
}

func PostUpdate(c *gin.Context){

	// get id from url
	id:= c.Param("id")

	//get data off the req body

	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)

	// find the post we should update
	var post models.Post
	initalizers.DB.First(&post,id)

	//update it
	initalizers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,Body: body.Body,
	})

	// respond with it
	c.JSON(200,gin.H{
		"post":post,
	})
}

func PostDelete(c *gin.Context){
	//get the id 


	// Delete the post


	// respond
	
}