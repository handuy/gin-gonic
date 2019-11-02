package main

import (
	"gin-gonic/jwt-cookie/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")

	router.GET("/", controller.Home)
	router.GET("/login", controller.LoginPage)
	router.POST("/login", controller.Login)

	router.GET("/secret", controller.CheckToken, controller.ViewSecret)

	router.Run(":8085")
}
