package main

import (
	"log"
	"time"

	model "gin-gonic/gorm/update-orm/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	ID        string       `gorm:"primary_key"`
	Name      string       `gorm:"type:varchar(50)"`
	Email     string       `gorm:"type:varchar(100)"`
	Age       int          `gorm:"type:BIGINT"`
	IsActive  bool      
	Average   float32      `gorm:"type:DECIMAL(6,2)"`
	CreatedAt time.Time
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	// Trước khi update
	var postInfo Post
	// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
	errGetPost := db.Where("id = ?", "123abc").Find(&postInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetPost != nil {
		log.Println(errGetPost)
		return
	}
	log.Println("Before update", postInfo)

	// Update 2 cột name và email trong bảng users
	postInfo.Name = "Kubernetes"
	postInfo.Email = "kuber@open.com"
	errUpdatePost := db.Save(&postInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errUpdatePost != nil {
		log.Println(errUpdatePost)
		return
	}

	// Kiểm tra xem đã update chưa
	var postInfoAfter Post
	// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
	errGetPost = db.Where("id = ?", "123abc").Find(&postInfoAfter).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetPost != nil {
		log.Println(errGetPost)
		return
	}
	log.Println("After update", postInfoAfter)

	defer db.Close()
}
