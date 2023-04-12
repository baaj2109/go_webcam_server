package global

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	Db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to create database, error: ", err)
		return
	}

	sqlDB, err := Db.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
