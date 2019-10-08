package main

import (
	model "gin-gonic/gorm/raw-update/model"
	// "log"
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
		DeptNo:   "d012",
		DeptName: "Technical",
	}

	tx := db.Begin()
	if err := tx.Error; err != nil {
		return
	}

	errorInsert := tx.Exec(`
		INSERT INTO departments VALUES (?, ?)`, department.DeptNo, department.DeptName).Error
	if errorInsert != nil {
		tx.Rollback()
		return
	}

	errorUpdate := tx.Exec(`
		UPDATEEEEE departments
		SET dept_name = 'Quality Test'
		WHERE dept_no = ?
	`, "d006").Error
	if errorUpdate != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
}
