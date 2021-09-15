package main

import (
	"gorm.io/gorm"
	"kalika-be/config"
	"kalika-be/database"
	"kalika-be/router"
	"sync"
)

var (
	syncOnce sync.Once
	dbh      *gorm.DB
)

func main()  {
	syncOnce.Do(DBConnection)
	dbConn, _ := dbh.DB()
	defer dbConn.Close()

	r := router.Routes(dbh)

	err := r.Start(":" + config.Get("APP_PORT").String())
	if err != nil {
		r.Logger.Fatal(err)
	}
}

func DBConnection() {
	dbh = database.Connect()
	database.Migrate()
}
