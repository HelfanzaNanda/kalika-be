package database

import (
	"fmt"
	"kalika-be/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() *gorm.DB {
	host := config.Get("DB_HOST").String()
	port := config.Get("DB_PORT").String()
	name := config.Get("DB_NAME").String()
	user := config.Get("DB_USER").String()
	password := config.Get("DB_PASSWORD").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
	dbConnPool, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if config.Get("DB_IS_DEBUG").Bool() {
		dbConnPool = dbConnPool.Debug()
	}

	sqlDB, err := dbConnPool.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxOpenConns(config.Get("DB_MAX_OPEN_CONNS").Int())
	sqlDB.SetMaxIdleConns(config.Get("DB_MAX_OPEN_CONNS").Int())
	sqlDB.SetConnMaxLifetime(config.Get("DB_CONN_MAX_LIFETIME").Duration())

	db = dbConnPool
	return db
}