package main

import (
	"log"
	"time"

	model "gin-gonic/gorm/delete-orm/model"

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

	// Trước khi delete
	var postInfo Post
	
	// WARNING !!!!!!!
	// Khi code API thì ID của post sẽ lấy từ đường dẫn hoặc form-data/JSON
	// Cần kiểm tra kĩ, không để xảy ra trường hợp postInfo.ID = ""
	// Nếu không gorm sẽ xoá sạch dữ liệu trong bảng --> DISASTER !!!!!
	postInfo.ID = "123abc"
	
	errDeletePost := db.Delete(&postInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errDeletePost != nil {
		log.Println(errDeletePost)
		return
	}

	// Kiểm tra xem đã delete chưa
	var postInfoAfter Post
	// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
	errGetPost := db.Where("id = ?", "123abc").Find(&postInfoAfter).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetPost != nil {
		log.Println(errGetPost)
		return
	}
	log.Println("After update", postInfoAfter)

	defer db.Close()
}
