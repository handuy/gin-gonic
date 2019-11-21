package main

import (
	"log"
	"time"

	model "gin-gonic/gorm/join/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	UserId     string
	UserName   string
	CardNumber string
	ExpiredAt  time.Time
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	// Dùng phpMyAdmin để INSERT dữ liệu mẫu vào 2 bảng users và credit_cards
	// insert into users values ("123", "Golang", 10)
	// insert into credit_cards VALUES ("abc", "2008-06-01 13:06:01", "123"), ("def", "2009-03-23 10:14:01", "123")

	var userInfo []UserInfo
	// Sử dụng column alias để cho tên bảng khớp với tên trường của struct hứng dữ liệu
	errGetUser := db.Table("users").
		Joins("join credit_cards on users.id = credit_cards.user_id").
		Select("users.id AS user_id, users.name AS user_name, credit_cards.number AS card_number, credit_cards.expired_at AS expired_at").
		Scan(&userInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}

	log.Println("User info", userInfo)

	defer db.Close()
}
