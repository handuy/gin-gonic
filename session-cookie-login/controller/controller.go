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
	switch session.Get("userId").(type) {
	case nil:
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Authenticated": false,
		})
	case string: 
		userName := session.Get("userEmail").(string)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "Welcome " + " " + userName,
			"Authenticated": true,
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
			Message: "Số điện thoại không được để trống",
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
				session.Set("userId", user.Id)
				session.Set("userEmail", user.Email)
				session.Save()

				c.JSON(http.StatusOK, LoginResponse{
					Success: true,
				})
				return
			}
		}

	}

	c.JSON(http.StatusUnauthorized, LoginResponse{
		Success: false,
	})
	return
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