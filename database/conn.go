package database

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *sql.DB
var err error
var Instance *gorm.DB

// func Connect(connectionString string) {
// 	Db, err = sql.Open("mysql", connectionString)
// 	if err != nil {
// 		log.Fatal(err)
// 		panic("Cannot connect to DB")
// 	}
// 	log.Println("Connected to Database...")
// }

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
		log.Println("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

// func Migrate() {
// 	Instance.AutoMigrate(&User{})
// 	log.Println("Database Migration Completed...")
// }
