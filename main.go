package main

import (
	"github.com/gin-gonic/gin"

	"gin-playground/controllers"
	"gin-playground/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.Run()
}
