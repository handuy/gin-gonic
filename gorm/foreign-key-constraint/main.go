package main

import (
	model "gin-gonic/gorm/raw-update/model"
	"log"
	"time"
)

type User struct {
	ID          string `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(100)"`
	Age         int
}

type CreditCard struct {
	Number    string `gorm:"primary_key"`
	ExpiredAt time.Time
	UserID    string
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var user User
	var creditCard CreditCard

	errCreateUser := db.CreateTable(user).Error
	if errCreateUser != nil {
		log.Println(errCreateUser)
		return
	}

	errCreateCard := db.CreateTable(creditCard).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error
	if errCreateCard != nil {
		log.Println(errCreateCard)
		return
	}

}
