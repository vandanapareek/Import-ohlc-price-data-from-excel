package controllers

import "go-price-data/services/csvdata"

var (
	csvService csvdata.ICsvService = csvdata.ICsvService(csvdata.CsvStruct{})
)

type Response struct {
	Msg string `json:"msg"`
}
