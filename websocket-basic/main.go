package main

import (
	"github.com/gin-gonic/gin"
	"gin-gonic/websocket-basic/controller"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("view/*")

    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{})
	})
	
	router.GET("/ws", func(c *gin.Context) {
		controller.Wshandler(c.Writer, c.Request)
	})

    router.Run("localhost:8087")
}