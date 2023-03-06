package main

import (
	"log"
	"net/http"

	"go-price-data/controllers"
	"go-price-data/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//https://stackoverflow.com/questions/52154609/fastest-way-of-reading-huge-files-in-go-with-small-ram

	//https://github.com/icodestuff-io/golang-docker-tutorial

	database.Connect("tester:secret@tcp(test_db_new)/test")

	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/search", controllers.Search)
	http.HandleFunc("/read-csv", controllers.ReadCsv)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
