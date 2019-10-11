package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/rs/xid"
)

type ErrorMesssage struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Success bool `json:"success"`
}

type Account struct {
	Id       string  `json:"id"`
	Name string  `json:"name"`
	Phone    string  `json:"phone"`
	Avatar   string  `json:"avatar"`
	Address  string  `json:"address"`
	Password string  `json:"password"`
	Message  string `json:"message"`
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
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	if phone == "" {
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
		Id: userID,
		Name: "",
		Phone:    phone,
		Avatar: "",
		Address: "",
		Password: passwordHash,
		Message: "success",
	}
	userList = append(userList, newUser)

	var userJSON = Account{
		Id: userID,
		Name: "",
		Phone:    phone,
		Avatar: "",
		Address: "",
		Message: "success",
	}
	c.JSON(http.StatusOK, userJSON)
	return
}

func Login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	for _, user := range userList {
		if user.Phone == phone {
			byteHash := []byte(user.Password)
			err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
			if err == nil {
				c.JSON(http.StatusOK, user)
				return
			}
		}

	}

	c.JSON(http.StatusUnauthorized, Account{
		Message: "fail",
	})
	return
}

func LogUsers(c *gin.Context) {
	log.Println("Current user list:", userList)
	c.Next()
}
