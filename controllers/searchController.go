package controllers

import (
	"encoding/json"
	"go-docker-tutorial/database"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

func Search(w http.ResponseWriter, r *http.Request) {
	filterStruct := database.SetDefault()

	err := schema.NewDecoder().Decode(&filterStruct, r.URL.Query())
	if err != nil {
		log.Println("Error in GET parameters : ", err)
	} else {
		log.Println("GET parameters : ", filterStruct)
	}
	var pd database.PriceData
	res := database.GetAllDetails(&pd, &filterStruct)
	if res == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Not found")
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
