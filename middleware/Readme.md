## Sử dụng middleware

1. Đăng kí hàm xử lý request
```go
func logTime(c *gin.Context) {
    log.Println("Current time:", time.Now())
    // Gọi hàm xử lý tiếp theo trong chuỗi middleware
	c.Next()
}
```

2. Đăng kí middleware cho từng route
```go
router := gin.Default()
router.LoadHTMLGlob("view/*.html")

router.GET("/", homePage)

// Đầu tiên chạy hàm logTime, trong logTime gọi c.Next() để gọi tiếp đến homePage
router.GET("/about", logTime, aboutPage)
```