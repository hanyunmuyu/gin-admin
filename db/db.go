package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
		DB = db
		sqlDB, _ := DB.DB()
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
	} else {
		fmt.Println(err)
	}
}
