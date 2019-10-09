package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func aboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}

func logTime(c *gin.Context) {
	log.Println("Current time:", time.Now())
	c.Next()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")

	router.GET("/", homePage)
	router.GET("/about", logTime, aboutPage)

	router.Run(":8085")
}
