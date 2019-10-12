package main

import (
	"gin-gonic/session-cookie-login/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"encoding/gob"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	// router.Use(logStore)
	// router.Use(sessions.Sessions("mysession", store))

	// Tạo cookie store dùng để lưu session bên trong secure cookie
	store := cookie.NewStore([]byte("tuananh"))
	store.Options(sessions.Options{
		HttpOnly: true,
	})
	gob.Register(&controller.Account{})

	// router.Use(sessions.Sessions("khapxungquanh", store))

	router.GET("/", sessions.Sessions("hungdung", store), controller.HomePage)
	router.GET("/signup", controller.SignupPage)
	router.GET("/login", controller.LoginPage)

	router.POST("/signup", controller.Register)
	router.POST("/login", sessions.Sessions("hungdung", store), controller.Login)
	router.GET("/logout", sessions.Sessions("hungdung", store), controller.Logout)

	// router.GET("/count", sessions.Sessions("hungdung", store), controller.AboutPage)

	router.Run(":8085")
}
