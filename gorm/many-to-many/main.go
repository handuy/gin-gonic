package main

import (
	model "gin-gonic/gorm/raw-update/model"
)

type Student struct {
	Id      string `gorm:"primary_key"`
	Name    string `gorm:"type:varchar(100)"`
	Age     int
	Classes []Class `gorm:"many2many:StudentClass;ASSOCIATION_JOINTABLE_FOREIGNKEY:IdClass;JOINTABLE_FOREIGNKEY:IdStudent"`
}

type Class struct {
	Id     string `gorm:"primary_key"`
	Name   string
	UserId string
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Class{})
	
	// .AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
