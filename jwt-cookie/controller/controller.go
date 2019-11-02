package controller

import (
	"log"
	"net/http"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Account struct {
	Email    string
	Password string
}

var userList = []Account{
	{
		Email:    "ronaldo@juve.com",
		Password: "juve",
	},
	{
		Email:    "messi@barca.com",
		Password: "barca",
	},
	{
		Email:    "neymar@psg.com",
		Password: "psg",
	},
}

const secretKey = "awesomejwt"

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	for _, user := range userList {
		if user.Email == email && user.Password == password {
			// Tạo token với Header lưu thông tin chung:
			// Loại token: JWT
			// Thuật toán mã hoá: HS256
			token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))

			// Truyền dữ liệu vào phần Claim của token
			// Dữ liệu có kiểu map[string]interface{} mô phỏng một cấu trúc dạng JSON
			token.Claims = jwt_lib.MapClaims{
				"Email": user.Email,
				"exp":   time.Now().Add(time.Hour * 1).Unix(),
			}

			// Tạo Signature cho token
			// Signature = HS256(Header, Claim, mysupersecretpassword)
			// Sử dụng mysupersecretpassword như một input đầu vào
			// để thuật toán HS256 tạo ra chuỗi signature
			tokenString, err := token.SignedString([]byte(secretKey))
			if err != nil {
				c.JSON(500, gin.H{"message": "Could not generate token"})
			}

			c.SetCookie("token", tokenString, 86400, "", "", false, true)
			c.JSON(200, "Thành công")
			return
		}
	}

	c.Redirect(http.StatusFound, "/fail")
	return
}

func CheckToken(c *gin.Context) {
	cookieValue, err := c.Cookie("token")
	if err != nil {
		log.Println("errir 0000", err)
		return
	}

	log.Println("cookie value", cookieValue)
	claims := jwt_lib.MapClaims{}

	tkn, err := jwt_lib.ParseWithClaims(cookieValue, claims, func(token *jwt_lib.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		if err == jwt_lib.ErrSignatureInvalid {
			log.Println("error 1")
			c.JSON(http.StatusUnauthorized, "Token không hợp lệ")
			return
		}
		log.Println("error 2", err)
		c.JSON(http.StatusBadRequest, "Request lỗi")
		return
	}

	if !tkn.Valid {
		log.Println("error 3")
		c.JSON(http.StatusUnauthorized, "Token không hợp lệ")
		return
	}

	log.Println("error 4")

	c.Next()
}

func ViewSecret(c *gin.Context) {
	c.HTML(http.StatusOK, "secret.html", gin.H{})
}
