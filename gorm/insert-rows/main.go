package main

import (
	"log"
	"time"

	model "gin-gonic/gorm/insert-rows/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Email     string
	Age       int
	IsActive  bool
	Average   float32
	CreatedAt time.Time
}

type Post struct {
	ID        string `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(50)"`
	Email     string `gorm:"type:varchar(100)"`
	Age       int    `gorm:"type:BIGINT"`
	IsActive  bool
	Average   float32 `gorm:"type:DECIMAL(6,2)"`
	CreatedAt time.Time
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	// Khi code API thì thông tin về user mới sẽ lấy từ file JSON hoặc form data của client gửi lên
	var newUser = User{
		Name:      "Golang",
		Email:     "golang@goole.com",
		Age:       10,
		IsActive:  true,
		Average:   8.64,
		CreatedAt: time.Now(),
	}
	errCreateUser := db.Create(&newUser).Error
	// Khi code API thì chỗ này trả về status 500 InternalServerError
	if errCreateUser != nil {
		log.Println(errCreateUser)
		return
	}

	// Khi code API thì thông tin về post mới sẽ lấy từ file JSON hoặc form data của client gửi lên
	var newPost = Post{
		ID: "123abc",
		Name:      "Golang",
		Email:     "golang@goole.com",
		Age:       10,
		IsActive:  true,
		Average:   8.64,
		CreatedAt: time.Now(),
	}
	errCreatePost := db.Create(&newPost).Error
	// Khi code API thì chỗ này trả về status 500 InternalServerError
	if errCreatePost != nil {
		log.Println(errCreatePost)
		return
	}

	defer db.Close()
}
