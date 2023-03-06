package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error
var Instance *gorm.DB

func Connect() {
	Instance, err = gorm.Open(mysql.Open("tester:secret@tcp(test_db_new)/test"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
		log.Println("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}
