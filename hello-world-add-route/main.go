package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	router.GET("/about", func(c *gin.Context) {
		c.String(200, "This is the about page")
	})
	router.GET("/khoa-hoc/:id/:name", func(c *gin.Context) {
		courseId := c.Param("id")
		courseName := c.Param("name")
		courseInfo := fmt.Sprintf("Thông tin khoá học: id khoá học là %s và tên khoá học là %s", courseId, courseName)

		c.String(200, courseInfo)
	})
	router.Run(":8080")
}
