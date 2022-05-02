package app

import (
	"fmt"

	"github.com/KwokBy/easy-ops/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGormMySql() *gorm.DB {
	config := configs.New()
	// 判断策略是否存在
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&&timeout=30s",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
