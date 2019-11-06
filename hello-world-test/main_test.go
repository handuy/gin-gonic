package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type About struct {
	Name string
}

type Contact struct {
	Email string
	Phone string
}

type Signup struct {
	Email string
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

// So sánh struct
func TestAboutRoute(t *testing.T) {
	router := gin.Default()
	router.GET("/about", func(c *gin.Context) {
		var info = About{
			Name: "Ngolo Kante",
		}
		c.JSON(200, info)
	})

	// Gửi GET request /about
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/about", nil)
	router.ServeHTTP(w, req)

	// Expected result
	var infoTest = About{
		Name: "Ngolo Kante",
	}

	// Lấy kết quả trả về từ GET /about và lưu vào biến target
	var target About
	json.NewDecoder(w.Result().Body).Decode(&target)
	log.Println("------------------", target)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, infoTest, target)
}

// So sánh slice
func TestContactRoute(t *testing.T) {
	router := gin.Default()
	router.GET("/contact", func(c *gin.Context) {
		var contact = []Contact{
			Contact{
				Email: "mbappe@gmail.com",
				Phone: "0123456789",
			},
			Contact{
				Email: "pogba@yahoo.com",
				Phone: "0239841485",
			},
			Contact{
				Email: "messi@barca.com",
				Phone: "0859347294",
			},
		}
		c.JSON(200, contact)
	})

	// Gửi GET request /contact
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/contact", nil)
	router.ServeHTTP(w, req)

	// Expected result
	var infoTest = []Contact{
		Contact{
			Email: "mbappe@gmail.com",
			Phone: "0123456789",
		},
		Contact{
			Email: "pogba@yahoo.com",
			Phone: "0239841485",
		},
		Contact{
			Email: "messi@barca.com",
			Phone: "0859347294",
		},
	}

	// Lấy kết quả trả về từ GET /about và lưu vào biến target
	var target []Contact
	json.NewDecoder(w.Result().Body).Decode(&target)
	log.Println("------------------", target)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, infoTest, target)
}

// Test POST form data
func TestFormRoute(t *testing.T) {
	router := gin.Default()
	router.POST("/signup", func(c *gin.Context) {
		email := c.PostForm("email")

		if email == "" {
			t.Fatal("Email không được để trống")
		}

		log.Println("Email received", email)
	})

	payload := url.Values{}
	payload.Set("email", "ronaldo@juve.com")

	// Gửi GET request /contact
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(payload.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
}

// Test POST upload binary file
func TestUploadRoute(t *testing.T) {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			t.Fatal(err)
			return
		}

		err = c.SaveUploadedFile(file, path.Join("./upload", file.Filename))
		if err != nil {
			t.Fatal(err)
			return
		}
	})

	file, err := os.Open("./git.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(part, file)
	writer.Close()

	// Gửi POST request /contact
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/upload", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	router.ServeHTTP(w, req)
}
