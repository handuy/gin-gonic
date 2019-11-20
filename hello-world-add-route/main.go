package main

import (
	"github.com/gin-gonic/gin"
)

type Issue struct {
	ID string
	Title string
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.GET("/issues", func(c *gin.Context) {
		var result = []Issue{
			{
				ID: "1",
				Title: "Issue số 1",
			},
			{
				ID: "2",
				Title: "Issue số 2",
			},
		}

		c.JSON(200, result)
	})

	router.GET("/issues/:id", func(c *gin.Context) {
		// Giả sử đây là danh sách các issue lưu trong database
		var listIssues = []Issue{
			{
				ID: "1",
				Title: "Issue số 1",
			},
			{
				ID: "2",
				Title: "Issue số 2",
			},
		}

		// result là kết quả trả về client
		var result = Issue{
			ID: "",
			Title: "",
		}

		// Lấy ID của issue mà người dùng muốn xem
		var issueId = c.Param("id")

		// Giả lập việc kiểm tra trong database xem có issue đó ko
		// Sau này khi kết nối đến database thì sẽ dùng lệnh SELECT để kiểm tra
		for _, v := range listIssues{
			if v.ID == issueId {
				result = v
				break
			}
		}

		c.JSON(200, result)
	})

	router.Run(":8081")
}
