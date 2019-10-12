package main

import (
	"gin-gonic/swagger/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	
	_ "gin-gonic/swagger/docs"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")

	url := ginSwagger.URL("http://localhost:8085/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.POST("/signup", controller.Register)
	router.POST("/login", controller.Login)

	router.Run(":8085")
}
