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

func registerPage(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{})
}

func getFormData(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	log.Println("Email:", email)
	log.Println("Password:", password)
}

func logTime(c *gin.Context) {
	log.Println("Current time:", time.Now())
	c.Next()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	router.Use(logTime)

	router.GET("/", homePage)
	router.GET("/about", aboutPage)
	router.GET("/form", registerPage)
	router.POST("/register", getFormData)

	router.Run(":8085")
}
