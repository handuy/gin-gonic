package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(50)"`
	Email     string    `gorm:"type:varchar(100)"`
	Age       int       `gorm:"type:BIGINT"`
	IsActive  bool      
	Average   float32   `gorm:"type:DECIMAL(6,2)"`
	CreatedAt time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost:8080)/employees?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	var user User
	var post Post

	db.CreateTable(user)
	db.CreateTable(post)

	defer db.Close()
}
