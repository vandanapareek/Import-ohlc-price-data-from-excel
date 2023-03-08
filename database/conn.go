package database

import (
	"database/sql"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error
var Instance *gorm.DB
var db *sql.DB

func ConnectToGorm() {
	Instance, err = gorm.Open(mysql.Open(os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_NAME")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
		log.Println("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Connect() {

	db, err = sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
		log.Println("Cannot connect to DB")
	}

	log.Println("Connected to Database...")
}
