package main

import (
	"log"
	"time"

	model "gin-gonic/gorm/update-orm-1/model"

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
	// Khi code API thì thông tin về post ID 
	// và các trường cần update cùng với giá trị tương ứng
	// sẽ lấy từ form-data/JSON (POST/PUT)
	// trong ví dụ này tạm thời fix cứng
	errUpdatePost := db.Model(&postInfo).Where("id = ?", "123abc").
		Updates(map[string]interface{}{"name":"AmazonWS", "email":"cloud@aws.com"}).
		Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errUpdatePost != nil {
		log.Println(errUpdatePost)
		return
	}

	// Kiểm tra xem đã update chưa
	var postInfoAfter Post
	// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request)
	errGetPost := db.Where("id = ?", "123abc").Find(&postInfoAfter).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetPost != nil {
		log.Println(errGetPost)
		return
	}
	log.Println("After update", postInfoAfter)

	defer db.Close()
}
