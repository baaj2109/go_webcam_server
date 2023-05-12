package global

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/baaj2109/webcam_server/settings"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite" // Sqlite driver based on GO
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	SQLLiteDb *gorm.DB
	MySQLDb   *sql.DB
	RedisDb   *redis.Client
)

func InitSQLLiteDb() {
	SQLLiteDb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to create database, error: ", err)
		return
	}

	sqlDB, err := SQLLiteDb.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

func InitMySqlDb(cfg *settings.MySqlConfig) error {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("failed to create database, error: ", err)
		return err
	}
	db.Logger.LogMode(1)
	sql, err := db.DB()
	if err != nil {
		panic(err)
	}
	sql.SetMaxIdleConns(cfg.MaxIdleConns)
	sql.SetMaxOpenConns(cfg.MaxOpenConns)
	sql.SetConnMaxLifetime(time.Second * time.Duration(cfg.MaxLifeTime))
	MySQLDb = sql
	return nil
}

func InitRedisDb(cfg *settings.RedisConfig) error {
	RedisDb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})
	result := RedisDb.Ping(context.Background())
	fmt.Println("redis ping:", result.Val())
	if result.Val() != "PONG" {
		// 连接有问题
		return result.Err()
	}
	return nil
}
