package main

import "github.com/gin-gonic/gin"

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
	router.Run(":8081")
}