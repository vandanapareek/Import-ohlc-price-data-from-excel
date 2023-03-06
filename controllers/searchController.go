package controllers

import (
	"encoding/json"
	"go-price-data/dto"
	"net/http"

	"github.com/gorilla/schema"
)

func Search(w http.ResponseWriter, r *http.Request) {
	//set default filters
	filterStruct := dto.SetDefault()

	//validate api params
	err := schema.NewDecoder().Decode(&filterStruct, r.URL.Query())
	if err != nil {
		var resp Response
		resp.Msg = "Error in GET parameters"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(resp)
		return
	}

	//call search service
	details, err := csvService.SearchDetails(filterStruct)
	if err != nil {
		//something went wrong while searching
		var resp Response
		resp.Msg = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(resp)
		return
	}
	var resp Response
	if len(details) == 0 || details == nil {
		resp.Msg = "No record found"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return

}

func HomePage(w http.ResponseWriter, req *http.Request) {
	var resp Response
	resp.Msg = "homepage"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
	return
}
