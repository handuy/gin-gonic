## Sử dụng middleware global

1. Đăng kí hàm xử lý request
```go
func logTime(c *gin.Context) {
    log.Println("Current time:", time.Now())
    // Gọi hàm xử lý tiếp theo trong chuỗi middleware
	c.Next()
}
```

2. Đăng kí middleware cho tất cả các route
```go
router := gin.Default()
router.LoadHTMLGlob("view/*.html")
router.Use(logTime)

router.GET("/", homePage)
router.GET("/about", aboutPage)
```