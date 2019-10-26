package main

import (
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
)

var (
	mysupersecretpassword = "unicornsAreAwesome"
)

func main() {
	r := gin.Default()

	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		// Tạo token với Header lưu thông tin chung:
		// Loại token: JWT
		// Thuật toán mã hoá: HS256
		token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))

		// Truyền dữ liệu vào phần Claim của token
		// Dữ liệu có kiểu map[string]interface{} mô phỏng một cấu trúc dạng JSON
		token.Claims = jwt_lib.MapClaims{
			"Id":  "Christopher",
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		}

		// Tạo Signature cho token
		// Signature = HS256(Header, Claim, mysupersecretpassword)
		// Sử dụng mysupersecretpassword như một input đầu vào 
		// để thuật toán HS256 tạo ra chuỗi signature 
		tokenString, err := token.SignedString([]byte(mysupersecretpassword))
		if err != nil {
			c.JSON(500, gin.H{"message": "Could not generate token"})
		}
		
		
		c.JSON(200, gin.H{"token": tokenString})
	})

	private := r.Group("/api/private")

	// Request gửi đến /api/private cần set Header: Authorization Bearer 
	private.Use(jwt.Auth(mysupersecretpassword))
	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from private"})
	})

	r.Run("localhost:8080")
}