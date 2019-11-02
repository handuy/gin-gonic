package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id       int    `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
	Phone    string `gorm:"type:varchar(100)"`
	Avatar   string `gorm:"type:varchar(100)"`
	Address  string `gorm:"type:varchar(100)"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost:8080)/employees?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	var user User

	db.CreateTable(user)

	defer db.Close()
}
