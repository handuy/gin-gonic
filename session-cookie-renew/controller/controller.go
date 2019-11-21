package controller

import (
	"net/http"
	"log"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AboutPage(c *gin.Context) {
	// var count int
	var startTime time.Time

	session := sessions.Default(c)

	switch v := session.Get("startTime").(type) {
	case nil:
		startTime = time.Now()
		session.Set("startTime", startTime)
		session.Save()
		log.Println("Begin set start time", startTime)
	case time.Time:
		startTime = v
		log.Println("start time", startTime)
		session.Save()
	}

	// switch v := session.Get("count").(type) {
	// case nil:
	// 	count = 0
	// 	session.Set("count", count)
	// 	session.Save()
	// case int:
	// 	count = v
	// 	count++
	// 	session.Set("count", count)
	// 	session.Save()
	// }

	c.JSON(http.StatusOK, startTime)
}
