package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
		DB = db
	} else {
		fmt.Println(err)
	}
}
