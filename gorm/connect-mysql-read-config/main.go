package main

import (
	model "gin-gonic/gorm/connect-mysql-read-config/model"
)
  
func main() {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
}