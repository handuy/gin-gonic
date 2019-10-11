package main

import (
	"gin-gonic/api-quynh/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	router.Use(controller.LogUsers)

	router.POST("/signup", controller.Register)
	router.POST("/login", controller.Login)

	router.Run(":8085")
}
