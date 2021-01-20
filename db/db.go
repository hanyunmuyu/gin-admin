package db

import (
	"fmt"
	"gin-admin/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB
)

func init() {
	v := utils.Config()
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if _, err := connect(); err != nil {
			fmt.Println(err)
		}
	})
	if _, err := connect(); err != nil {
		fmt.Println(err)
	}
}
func connect() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		utils.Config().GetString("mysql.username"),
		utils.Config().GetString("mysql.password"),
		utils.Config().GetString("mysql.host"),
		utils.Config().GetString("mysql.port"),
		utils.Config().GetString("mysql.db"),
		utils.Config().GetString("mysql.charset"),
	)
	//dsn := "root:123456@tcp(127.0.0.1:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
		DB = db
		sqlDB, _ := DB.DB()
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
		return DB, nil
	} else {
		return nil, err
	}
}
