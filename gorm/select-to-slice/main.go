package main

import (
	"log"
	"time"

	model "gin-gonic/gorm/select-to-slice/model"

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

	var userInfo []User
	errGetUser := db.Find(&userInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}

	for _,v := range userInfo {
		log.Println("Employee", v)
	}

	defer db.Close()
}
