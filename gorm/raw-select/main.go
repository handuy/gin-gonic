package main

import (
	"log"
	model "gin-gonic/gorm/raw-select/model"
) 

type Employee struct {
	Name string
	Age  int
	Gender string
} 
  
func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var employeeList []Employee
	db.Raw(`
		SELECT CONCAT(first_name, ' ', last_name) AS name, 
		ROUND( DATEDIFF( CURRENT_DATE(), birth_date)/365, 0 ) AS age,
		gender
		FROM employees
		ORDER BY last_name DESC
		LIMIT 30
		OFFSET 0
	`).Scan(&employeeList)

	for _, employee := range employeeList {
        log.Println("Employee", employee)
    }
}