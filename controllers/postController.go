package controllers

import "github.com/gin-gonic/gin"

func PostCreate(c *gin.Context) {

	//get data from request body

	//create a post
	
	//return it
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
