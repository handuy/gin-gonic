package main

import (
	"gin-gonic/upload-file/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")

	router.GET("/", controller.HomePage)

	router.POST("/upload-video", controller.UploadVideo)
	router.Run(":8085")
}
