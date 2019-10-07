package main

import (
	model "gin-gonic/gorm/raw-update/model"
	"time"
)

type User struct {
	Id          string `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(100)"`
	Age         int
}

type CreditCard struct {
	Number    string `gorm:"primary_key"`
	ExpiredAt time.Time
	UserId    string
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
