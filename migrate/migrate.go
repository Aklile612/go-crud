package main

import (
	"github.com/Aklile612/go-crud/initalizers"
	"github.com/Aklile612/go-crud/models"
)

func init(){
	initalizers.LoadEnvVariables()
	initalizers.ConnectionToDB()
}

func main() {
	initalizers.DB.AutoMigrate(&models.Post{})
}