package app

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGormMySql() *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&&timeout=30s",
		"easy_ops",
		"Gl@987963951",
		"42.192.11.9",
		3306,
		"easy_ops",
	)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
