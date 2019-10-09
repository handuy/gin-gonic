## Upload form data

1. HTML form có dạng:
```html
<!-- Gửi POST request lên /register -->
<form action="/register" method="POST">
    <div class="container">
        <h1>Đăng kí tài khoản</h1>
        <hr>
    
        <label for="email"><b>Email</b></label>
        <input type="text" placeholder="Nhập Email" name="email" required>
    
        <label for="password"><b>Password</b></label>
        <input type="password" placeholder="Nhập Password" name="password" required>
    
        <hr>
    
        <button type="submit" class="registerbtn">Đăng kí</button>
    </div>
</form>
```

2. Đăng kí GET route trả về form đăng kí trên 
```go
func registerPage(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{})
}

router.GET("/form", registerPage)
```

3. Đăng kí POST route /register để nhận form data
```go
func getFormData(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	log.Println("Email:", email)
	log.Println("Password:", password)
}

router.POST("/register", getFormData)
```