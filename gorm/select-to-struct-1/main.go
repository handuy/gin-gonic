package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Employee struct {
	EmpNo      int
	BeginName  string
	FinalName   string
	Sex     string
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost:8080)/employees?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	var firstEmployee Employee
	db.Table("employees").Select("emp_no, first_name AS begin_name, last_name AS final_name, gender AS sex").Where("emp_no = ?", 10001).Scan(&firstEmployee)

	log.Println("Employee", firstEmployee)

	defer db.Close()
}
