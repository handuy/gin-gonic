## Sử dụng session cookie lưu thông tin đăng nhập

1. Xem trước ví dụ về session-cookie tại:
https://github.com/handuy/gin-gonic/tree/master/session-cookie

2. Khi đăng nhập thành công, lưu thông tin UserId và UserEmail vào session
```go
session := sessions.Default(c)
session.Set("userId", user.Id)
session.Set("userEmail", user.Email)
session.Save()
```

3. Khi gửi GET request đến **/**, kiểm tra thông tin session 
```go
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
```