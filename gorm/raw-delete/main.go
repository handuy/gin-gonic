package main

import (
	model "gin-gonic/gorm/raw-delete/model"
	"log"
)

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)
	
	errDelete := db.Exec(`
		DELETE FROM users
		WHERE id = ?
	`, 2).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errDelete != nil {
		log.Println(errDelete)
		return
	}
	
}
