package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Setup the Mysql database, the username with password
func Connect() {
	d, err := gorm.Open("mysql", "library_user:pass123@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// return the mysql database
func GetDB() *gorm.DB {
	return db
}
