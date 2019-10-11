package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	var count = 0
	c.JSON(http.StatusOK, count)
}

func AboutPage(c *gin.Context) {
	var count int

	session := sessions.Default(c)
	switch v := session.Get("count").(type) {
	case nil:
		count = 0
		session.Set("count", count)
		session.Save()
	case int: 
		count = v
		count++
		session.Set("count", count)
		session.Save()
	}

	c.JSON(http.StatusOK, count)
}
