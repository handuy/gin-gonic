package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type About struct {
	Name string
}

func TestHomeRoute(t *testing.T) {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello World", w.Body.String())
}

func TestAboutRoute(t *testing.T) {
	router := gin.Default()
	router.GET("/about", func(c *gin.Context) {
		var info = About{
			Name: "Ngolo Kante",
		}
		c.JSON(200, info)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/about", nil)
	router.ServeHTTP(w, req)

	var infoTest = About{
		Name: "Ngolo Kante",
	}

	var target About
	json.NewDecoder(w.Result().Body).Decode(&target)
	log.Println("------------------", target)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, infoTest, target)
}
