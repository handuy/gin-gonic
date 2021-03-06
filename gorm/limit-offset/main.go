package main

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Employee struct {
	EmpNo     int
	BirthDate time.Time
	FirstName string
	LastName  string
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost:8080)/employees?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	var employeeList []Employee
	orderBy := "last_name DESC"
	limit := 10
	offset := 5

	db.Table("employees").Select("emp_no, birth_date, first_name, last_name").Where("first_name LIKE ?", "Cristinel%").Order(orderBy).Limit(limit).Offset(offset).Scan(&employeeList)

	for _, employee := range employeeList {
        log.Println("Employee", employee)
    }

	defer db.Close()
}
