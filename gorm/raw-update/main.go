package main

import (
	model "gin-gonic/gorm/raw-update/model"
	"log"
)

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	rowAffected := db.Exec(`
		UPDATE employees
		SET first_name = 'Lionel', last_name = 'Messi', gender = 'M'
		WHERE emp_no = ?
	`, 10001).RowsAffected

	log.Println("row affected", rowAffected)
}
