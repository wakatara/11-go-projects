package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open("root:rory57tate@/bookstore?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
