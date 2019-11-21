package main

import (
	model "gin-gonic/gorm/transaction/model"
	"log"
	"time"
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

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	tx := db.Begin()

	// Khi code API thì thông tin về user mới sẽ lấy từ file JSON hoặc form data của client gửi lên
	var newUser = User{
		Name:      "PHP",
		Email:     "php@elephant.com",
		Age:       20,
		IsActive:  false,
		Average:   8.64,
		CreatedAt: time.Now(),
	}
	errCreateUser := tx.Create(&newUser).Error
	// Khi code API thì chỗ này trả về status 500 InternalServerError
	if errCreateUser != nil {
		log.Println(errCreateUser)
		tx.Rollback()
		return
	}

	// Trước khi update
	var userInfo User
	// Khi code API thì thông tin về user ID 
	// và các trường cần update cùng với giá trị tương ứng
	// sẽ lấy từ form-data/JSON (POST/PUT)
	// trong ví dụ này tạm thời fix cứng
	errUpdateUser := tx.Model(&userInfo).Where("id = ?", 3).
		Updates(map[string]interface{}{"name":"AmazonWS", "email":"cloud@aws.com"}).
		Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errUpdateUser != nil {
		log.Println(errUpdateUser)
		tx.Rollback()
		return
	}

	tx.Commit()
}
