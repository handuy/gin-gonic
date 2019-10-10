## Render HTML trả về client

1. Đăng kí thư mục chứa các file HTML
```go
router.LoadHTMLGlob("view/*")
```

2. Trong file HTML sẽ gọi đến các file CSS để style trang web. Các file CSS được để ở thư mục resources/css
```go
router.Static("/resources", "./resources")
```

3. GET request gửi đến **/** được xử lý bởi **getHomePage**: Truyền data vào file **index.html** rồi trả về client
```go
type Club struct {
	Name string
	Year int
	Nation string
}

func getHomePage(c *gin.Context) {
	var clubList = []Club{
		{
			Name: "Man Utd",
			Year: 1879,
			Nation: "England",
		},
		{
			Name: "Juventus",
			Year: 1894,
			Nation: "Italy",
		},
		{
			Name: "Bayern",
			Year: 1901,
			Nation: "Germany",
		},
	}

	var data = gin.H{
		"Name": "Juan Mata",
		"Age": 30,
		"CurrentClub": "Man Utd",
		"ClubList": clubList,
	}

	c.HTML(http.StatusOK, "index.html", data)
}

router.GET("/", getHomePage)
```

4. File **index.html** chứa các logic xử lý đơn giản: **If/Else**, chạy vòng lặp **range**, render dữ liệu struct. 
Tìm hiểu chi tiết hơn về HTML templates tại: https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet