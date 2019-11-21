package main

import (
	"gin-gonic/logging/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")

	router.GET("/", controller.HomePage)
	router.POST("/upload-video", controller.UploadVideo)
	router.Run(":8085")
}

// khai báo biến với các kiểu dữ liệu
var age int
var name string
var result bool
    
// gán giá trị cho biến
age = 26
name = "duy"
result = true