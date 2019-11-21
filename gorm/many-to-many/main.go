package main

import (
	model "gin-gonic/gorm/raw-update/model"
	"log"
)

type Student struct {
	Id   string `gorm:"primary_key"`
	Name string `gorm:"type:varchar(100)"`
}

type Class struct {
	Id   string `gorm:"primary_key"`
	Name string
}

type StudentClass struct {
	StudentId string `gorm:"primary_key"`
	ClassId   string `gorm:"primary_key"`
}

func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var student Student
	var class Class
	var studentClass StudentClass

	errCreateStudent := db.AutoMigrate(student).Error
	if errCreateStudent != nil {
		log.Println(errCreateStudent)
		return
	}

	errCreateClass := db.AutoMigrate(class).Error
	if errCreateClass != nil {
		log.Println(errCreateClass)
		return
	}

	errCreateStudentClass := db.AutoMigrate(studentClass).
				AddForeignKey("student_id", "students(id)", "RESTRICT", "RESTRICT").
				AddForeignKey("class_id", "classes(id)", "RESTRICT", "RESTRICT").Error
	if errCreateStudentClass != nil {
		log.Println(errCreateStudentClass)
		return
	}
}
