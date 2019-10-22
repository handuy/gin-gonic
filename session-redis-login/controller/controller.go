package controller

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type ErrorMesssage struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Success bool `json:"success"`
}

type Account struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var userList = []Account{}

func HomePage(c *gin.Context) {
	session := sessions.Default(c)
	v, ok := session.Get("user").(*Account)
	if ok {
		userName := v.Email
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "Welcome " + " " + userName,
			"Authenticated": true,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Authenticated": false,
		})
	}
}

func SignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "" {
		c.JSON(http.StatusBadRequest, ErrorMesssage{
			Message: "Email không được để trống",
		})
		return
	}

	if len(password) < 4 {
		c.JSON(http.StatusBadRequest, ErrorMesssage{
			Message: "Mật khẩu phải có tối thiểu 4 kí tự",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	passwordHash := string(hash)
	userID := xid.New().String()

	var newUser = Account{
		Id:       userID,
		Email:    email,
		Password: passwordHash,
	}
	userList = append(userList, newUser)

	log.Println("user list", userList)

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
				session := sessions.Default(c)
				// session.Set("userId", user.Id)
				// session.Set("userEmail", user.Email)
				log.Println("Jack Vietnam")
				session.Set("user", user)
				session.Save()

				c.Redirect(http.StatusFound, "/")
				return
			}
		}

	}

	c.Redirect(http.StatusFound, "/")
	return
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()

	c.Redirect(http.StatusFound, "/")
}

func AboutPage(c *gin.Context) {
	var count int

	session := sessions.Default(c)
	switch v := session.Get("count").(type) {
	case nil:
		count = 0
		session.Set("count", count)
		session.Save()
	case int: 
		count = v
		count++
		session.Set("count", count)
		session.Save()
	}

	c.JSON(http.StatusOK, count)
}