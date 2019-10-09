## Lưu mật khẩu người dùng dưới dạng mã băm bcrypt hashed

1. Khi user gửi form đăng kí, lưu password dưới dạng mã băm, sử dụng package bcrypt:
```go
import "golang.org/x/crypto/bcrypt"

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
```

2. Khi user gửi form đăng nhập, so sánh password gửi lên với hashed password bằng method CompareHashAndPassword 
```go
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
```