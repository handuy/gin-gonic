package main

import (
	model "gin-gonic/gorm/transaction-rollback/model"
	"log"
	"time"
	// "log"
)

type User struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Email     string
	Age       int
	IsActive  bool
	Average   float32
	CreatedAt time.Time
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	tx := db.Begin()
	if err := tx.Error; err != nil {
		return
	}

	// Khi code API thì thông tin về user mới sẽ lấy từ file JSON hoặc form data của client gửi lên
	var newUser = User{
		Name:      "Java",
		Email:     "java@oracle.com",
		Age:       20,
		IsActive:  true,
		Average:   8.64,
		CreatedAt: time.Now(),
	}
	errCreateUser := tx.Create(&newUser).Error
	// Khi code API thì chỗ này trả về status 500 InternalServerError
	if errCreateUser != nil {
		log.Println(errCreateUser)
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
