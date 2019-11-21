package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	model "gin-gonic/gorm/select-to-struct-other/model"
)

type UserInfo struct {
	UserId      int
	FullName   string
	Email    string
	ActiveOrNot   int
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var userInfo UserInfo
	// Sử dụng column alias để cho tên bảng khớp với tên trường của struct hứng dữ liệu
	errGetUser := db.Table("users").
				Select("id AS user_id, name AS full_name, email, is_active AS active_or_not").
				Where("id = ?", 2).Scan(&userInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}

	log.Println("Employee", userInfo)

	defer db.Close()
}
