package main

import (
	"gin-gonic/logging-db/controller"

	"github.com/gin-gonic/gin"

	model "gin-gonic/logging-db/model"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()

	var logInfo controller.ErrorInfo
	db.AutoMigrate(logInfo)

	control := &controller.Controller{DB: db}

	router.GET("/", control.HomePage)
	router.POST("/upload-video", control.UploadVideo)
	router.Run(":8085")
}
