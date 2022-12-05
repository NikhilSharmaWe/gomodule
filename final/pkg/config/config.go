package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:nikhil@tcp(127.0.0.1:3306)/test01?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}
