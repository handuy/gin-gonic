package main

import (
	"time"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	model "gin-gonic/gorm/create-table/model"
) 

type User struct {
	ID        int       `gorm:"primary_key"`
	Name      string    
	Email     string    
	Age       int       
	IsActive  bool      
	Average   float32   
	CreatedAt time.Time
}

type Post struct {
	ID        string       `gorm:"primary_key"`
	Name      string       `gorm:"type:varchar(50)"`
	Email     string       `gorm:"type:varchar(100)"`
	Age       int          `gorm:"type:BIGINT"`
	IsActive  bool      
	Average   float32      `gorm:"type:DECIMAL(6,2)"`
	CreatedAt time.Time
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var user User
	var post Post

	errCreateUser := db.CreateTable(user).Error
	if errCreateUser != nil {
		log.Println(errCreateUser)
		return
	}

	errCreatePost := db.CreateTable(post).Error
	if errCreatePost != nil {
		log.Println(errCreatePost)
		return
	}

	defer db.Close()
}
