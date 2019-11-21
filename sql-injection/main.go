package main

import (
	"fmt"
	model "gin-gonic/sql-injection/model"
	"log"
)

type Employee struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	injection := ";DROP DATABASE demo;"
	queryString := fmt.Sprintf(`
		SELECT * FROM employees %s
	`, injection)

	var employeeList []Employee
	db.Raw(queryString).Scan(&employeeList)

	for _, employee := range employeeList {
		log.Println("Employee", employee)
	}
}
