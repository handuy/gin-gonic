package main

import "github.com/gin-gonic/gin"

type userJSON struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func main() {
	router := gin.Default()
	router.PUT("/update-profile", func(c *gin.Context) {
		var userInfo userJSON
		err := c.BindJSON(&userInfo)
		if err != nil {
			return
		}

		// Logic update profile
	})
	router.DELETE("/delete-issue/:id", func(c *gin.Context) {
		issueId := c.Param("id")

		// Logic xoá issue
	})
	router.Run(":8080")
}