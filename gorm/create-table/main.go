package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	model "gin-gonic/gorm/create-table/model"
) 

type Company struct {
	Id        int `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(100)"`
	Address   string `gorm:"type:varchar(100)"`
	IsGlobal  bool
	CreatedAt time.Time
	OtherInfo model.ManagerInfo   `sql:"TYPE:json"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost:8080)/employees?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	var company Company

	db.CreateTable(company)

	defer db.Close()
}
