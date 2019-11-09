package main

import (
	"gin-gonic/api-long-flutter/controller"
	"gin-gonic/api-long-flutter/model"
	"github.com/gin-gonic/contrib/jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	ginController := controller.NewController()

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()

	ginController.DB = db
	ginController.Config = config

	router.POST("/signup", ginController.Register)
	router.POST("/register", ginController.RegisterJSON)

	// router.POST("/login", ginController.Login)
	router.POST("/login", ginController.LoginJSON)

	router.GET("/issues", jwt.Auth(model.SecretKey), ginController.ListIssues)
	router.GET("/issues/:id", jwt.Auth(model.SecretKey), ginController.IssueDetail)
	router.POST("/create-issue", jwt.Auth(model.SecretKey), ginController.CreateIssue)

	router.GET("/profile", jwt.Auth(model.SecretKey), ginController.ProfileDetail)

	router.POST("/upload-file", jwt.Auth(model.SecretKey), ginController.UploadFile)

	router.GET("/media/static/:userId/:fileId", jwt.Auth(model.SecretKey), ginController.ServeFile)

	router.Run(":8086")
}
