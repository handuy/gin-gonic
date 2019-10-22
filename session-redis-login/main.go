package main

import (
	"encoding/gob"
	"gin-gonic/session-redis-login/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	// router.Use(logStore)
	// router.Use(sessions.Sessions("mysession", store))

	// Kết nối redis database để lưu session ID
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "123", []byte("secret"))
	// store.Options(sessions.Options{
	// 	HttpOnly: true,
	// })
	gob.Register(&controller.Account{})
	// router.Use(sessions.Sessions("myredissession", store))

	// router.Use(sessions.Sessions("khapxungquanh", store))

	router.GET("/", sessions.Sessions("helloredis", store), controller.HomePage)
	router.GET("/signup", controller.SignupPage)
	router.GET("/login", controller.LoginPage)

	router.POST("/signup", controller.Register)
	router.POST("/login", sessions.Sessions("helloredis", store), controller.Login)
	router.GET("/logout", sessions.Sessions("helloredis", store), controller.Logout)

	// router.GET("/count", sessions.Sessions("hungdung", store), controller.AboutPage)

	router.Run(":8085")
}
