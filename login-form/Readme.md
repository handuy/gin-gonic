## Xử lý logic đăng nhập

1. HTML form có dạng:
```html
<!-- Gửi POST request lên /login -->
<form action="/login" method="POST">
    <div class="container">
        <h1>Đăng nhập</h1>
        <hr>
    
        <label for="email"><b>Email</b></label>
        <input type="text" placeholder="Nhập Email" name="email" required>
    
        <label for="password"><b>Password</b></label>
        <input type="password" placeholder="Nhập Password" name="password" required>
    
        <hr>
    
        <button type="submit" class="registerbtn">Đăng nhập</button>
    </div>
</form>
```

2. Đăng kí GET route trả về form đăng nhập 
```go
func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{})
}

router.GET("/login", loginPage)
```

3. Đăng kí POST route /login để nhận form data và check email+password
```go
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

router.POST("/login", getFormData)
```