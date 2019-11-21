package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	model "gin-gonic/gorm/select-to-struct/model"
)

type UserInfo struct {
	Id      int
	Name   string
	Email    string
	IsActive   int
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var userInfo UserInfo
	errGetUser := db.Table("users").Select("id, name, email, is_active").
				Where("id = ?", 2).Scan(&userInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}

	log.Println("Employee", userInfo)

	defer db.Close()
}
