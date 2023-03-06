package main

import (
	"log"
	"net/http"

	"go-price-data/controllers"
	"go-price-data/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.Connect()

	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/search", controllers.Search)
	http.HandleFunc("/read-csv", controllers.ReadCsv)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
