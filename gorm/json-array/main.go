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

	var managerInfo = model.ManagerInfo{
		Name: "Max Allegri",
		Age: 42,
		HireDate: time.Now(),
	}

	var company1 = Company{
		Id: 4,
		Name: "Juventus",
		Address: "Turin",
		IsGlobal: true,
		CreatedAt: time.Now(),
		OtherInfo: managerInfo,
	}

	var company2 = Company{
		Id: 5,
		Name: "Inter",
		Address: "Milan",
		IsGlobal: true,
		CreatedAt: time.Now(),
		OtherInfo: managerInfo,
	}

	var company = []Company{company1, company2}

	db.Create(company)

	defer db.Close()
}
