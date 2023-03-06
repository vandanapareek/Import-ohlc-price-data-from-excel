package controllers

import (
	"encoding/json"
	"go-price-data/dto"
	"go-price-data/errors"
	"net/http"

	"github.com/gorilla/schema"
)

func Search(w http.ResponseWriter, r *http.Request) {
	//set default filters
	filterStruct := dto.SetDefault()

	//validate api params
	err := schema.NewDecoder().Decode(&filterStruct, r.URL.Query())
	if err != nil {
		http.Error(w, errors.InvalidParamError.Message, http.StatusUnprocessableEntity)
		return
	}

	//call search service
	details, err := csvService.SearchDetails(filterStruct)
	if err != nil {
		//something went wrong while searching
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if len(details) == 0 || details == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("No record found")
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(details)
		return
	}

}

func HomePage(w http.ResponseWriter, req *http.Request) {
	var resp Response
	resp.Msg = "homepage"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
	return
}
