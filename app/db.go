package app

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGormMySql() *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&&timeout=30s",
		"root",
		"Gl987963951",
		"127.0.0.1",
		3306,
		"easy_ops",
	)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
