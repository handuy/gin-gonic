package main

import (
	"gin-gonic/session-cookie/controller"

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

	// router.Use(sessions.Sessions("khapxungquanh", store))

	router.GET("/", controller.HomePage)
	router.GET("/about", sessions.Sessions("hungdung", store), controller.AboutPage)

	router.Run(":8085")
}
