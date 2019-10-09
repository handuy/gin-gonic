package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Account struct {
	Email    string
	Password string
}

var userList = []Account{
	{
		Email: "ronaldo@juve.com",
		Password: "juve",
	},
	{
		Email: "messi@barca.com",
		Password: "barca",
	},
	{
		Email: "neymar@psg.com",
		Password: "psg",
	},
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func aboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}

func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{})
}

func successPage(c *gin.Context) {
	c.HTML(http.StatusOK, "success.html", gin.H{})
}

func failPage(c *gin.Context) {
	c.HTML(http.StatusOK, "fail.html", gin.H{})
}

func getFormData(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	for _, user := range userList {
		if user.Email == email && user.Password == password {
			c.Redirect(http.StatusFound, "/success")
			return
		}
	}

	c.Redirect(http.StatusFound, "/fail")
	return
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
	router.GET("/login", loginPage)
	router.POST("/login", getFormData)

	router.GET("/success", successPage)
	router.GET("/fail", failPage)

	router.Run(":8085")
}
