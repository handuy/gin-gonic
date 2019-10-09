package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Email    string
	Password string
}

var userList = []Account{}

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func AboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}

func SignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func SuccessPage(c *gin.Context) {
	c.HTML(http.StatusOK, "success.html", gin.H{})
}

func FailPage(c *gin.Context) {
	c.HTML(http.StatusOK, "fail.html", gin.H{})
}

func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	passwordHash := string(hash)

	var newUser = Account{
		Email:    email,
		Password: passwordHash,
	}
	userList = append(userList, newUser)

	c.Redirect(http.StatusFound, "/")
	return
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	for _, user := range userList {
		if user.Email == email {
			byteHash := []byte(user.Password)
			err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
			if err == nil {
				c.Redirect(http.StatusFound, "/success")
				return
			}
		}

	}

	c.Redirect(http.StatusFound, "/fail")
	return
}

func LogUsers(c *gin.Context) {
	log.Println("Current user list:", userList)
	c.Next()
}
