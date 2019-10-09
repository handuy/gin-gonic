package main

import (
	"gin-gonic/hash-password-bcrypt/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	router.Use(controller.LogUsers)

	router.GET("/", controller.HomePage)
	router.GET("/about", controller.AboutPage)

	router.GET("/signup", controller.SignupPage)
	router.POST("/signup", controller.Register)
	router.GET("/login", controller.LoginPage)
	router.POST("/login", controller.Login)

	router.GET("/success", controller.SuccessPage)
	router.GET("/fail", controller.FailPage)

	router.Run(":8085")
}
