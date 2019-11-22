package main

import (
	model "gin-gonic/gorm/raw-insert/model"
	"log"
	"time"
)

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	errInsert := db.Exec(`
		INSERT INTO users (name, email, age, is_active, average, created_at) VALUES 
		(?, ?, ?, ?, ?, ?)
	`, "Azure", "azure@microsoft.com", 11, true, 9.87, time.Now()).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errInsert != nil {
		log.Println(errInsert)
		return
	}

}
