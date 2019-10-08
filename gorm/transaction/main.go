package main

import (
	model "gin-gonic/gorm/raw-update/model"
)

type Department struct {
	DeptNo   string
	DeptName string
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var department = Department{
		DeptNo:   "d010",
		DeptName: "Accounting",
	}

	tx := db.Begin()

	tx.Create(department)

	tx.Exec(`
		UPDATE departments
		SET dept_name = 'Quality Control'
		WHERE dept_no = ?
	`, "d006")

	tx.Commit()
}
