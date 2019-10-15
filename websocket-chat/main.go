package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("view/*")

	go NewHub.Run()

    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{})
	})
	
	router.GET("/ws", func(c *gin.Context) {
		Wshandler(c.Writer, c.Request)
	})

    router.Run("localhost:8087")
}