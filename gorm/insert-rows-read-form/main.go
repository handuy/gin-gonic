package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
) 

type User struct {
	Id       string `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
	Phone    string `gorm:"type:varchar(100)"`
	Avatar   string `gorm:"type:varchar(100)"`
	Address  string `gorm:"type:varchar(100)"`
}

func main() {
	router := gin.Default()

	db, err := gorm.Open("mysql", "root:123@(localhost:8080)/employees?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	defer db.Close()

	router.POST("/register", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		var user = User{
			Id: xid.New().String(),
			Email: email,
			Password: password,
		}

		err := db.Create(user).Error
		if err != nil {
			c.String(http.StatusInternalServerError, "Server error")
			return
		}

		err = db.Exec(`INSERT INTO user(id, email, password) VALUES ?, ?, ?`, xid.New().String(), email, password).Error
		if err != nil {
			c.String(http.StatusInternalServerError, "Server error")
			return
		}
	})

	router.Run(":8085")
}
