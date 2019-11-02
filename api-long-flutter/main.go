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
	router.POST("/signup-json", ginController.RegisterJSON)

	router.POST("/login", ginController.Login)
	router.POST("/login-json", ginController.LoginJSON)

	router.GET("/issues", jwt.Auth(model.SecretKey), ginController.ListIssues)
	router.GET("/issues/:id", jwt.Auth(model.SecretKey), ginController.IssueDetail)
	router.POST("/create-issue", jwt.Auth(model.SecretKey), ginController.CreateIssue)

	router.GET("/profile", jwt.Auth(model.SecretKey), ginController.ProfileDetail)

	router.Run(":8085")
}
