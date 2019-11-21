package main

import (
    _ "github.com/go-sql-driver/mysql"
    "xorm.io/xorm"
)

var engine *xorm.Engine

func main() {
    var err error
	engine, err = xorm.NewEngine("mysql", "root:123@(localhost:8080)/employees?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}