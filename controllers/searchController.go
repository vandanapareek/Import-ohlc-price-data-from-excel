package controllers

import (
	"encoding/json"
	"go-price-data/database"
	"go-price-data/dto"
	"net/http"

	"github.com/gorilla/schema"
)

func Search(w http.ResponseWriter, r *http.Request) {
	//set default filters
	filterStruct := dto.SetDefault()

	//check
	err := schema.NewDecoder().Decode(&filterStruct, r.URL.Query())
	if err != nil {
		var resp Response
		resp.Code = 422
		resp.Msg = "Error in GET parameters"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(resp)
		return
	}
	var pd database.PriceData
	res := database.GetAllDetails(&pd, filterStruct)
	if res == nil {
		var resp Response
		resp.Code = 422
		resp.Msg = "No record found"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		var resp Response
		resp.Code = 200
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return
	}
}

func HomePage(w http.ResponseWriter, req *http.Request) {
	var resp Response
	resp.Code = 200
	resp.Msg = "homepage"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
	return
}
