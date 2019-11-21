package main

import (
	"log"
	"time"

	model "gin-gonic/gorm/update-orm/model"

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

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	// Trước khi update
	var userInfo User
	// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
	errGetUser := db.Where("id = ?", 3).Find(&userInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}
	log.Println("Before update", userInfo)

	// Update 2 cột name và email trong bảng users
	userInfo.Name = "Kubernetes"
	userInfo.Email = "kuber@open.com"
	errUpdateUser := db.Save(&userInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errUpdateUser != nil {
		log.Println(errUpdateUser)
		return
	}

	// Kiểm tra xem đã update chưa
	var userInfoAfter User
	// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
	errGetUser = db.Where("id = ?", 3).Find(&userInfoAfter).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}
	log.Println("After update", userInfoAfter)

	defer db.Close()
}
