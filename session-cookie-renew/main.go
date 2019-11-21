package main

import (
	"gin-gonic/session-cookie-renew/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	// router.Use(logStore)
	// router.Use(sessions.Sessions("mysession", store))

	// Tạo cookie store dùng để lưu session bên trong secure cookie
	store := cookie.NewStore([]byte("tuananh"))
	// store.Options(sessions.Options{
	// 	HttpOnly: true,
	// })

	// router.Use(sessions.Sessions("khapxungquanh", store))

	router.GET("/", sessions.Sessions("tuananh", store), controller.AboutPage)

	router.Run(":8089")
}
