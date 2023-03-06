package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go-price-data/consts"
	"go-price-data/database"
	"go-price-data/errors"
	"go-price-data/services/dataProcess"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	dataProcessService dataProcess.ICsvService = dataProcess.ICsvService(dataProcess.CsvStruct{})
)

func ReadCsv(w http.ResponseWriter, req *http.Request) {

	err := processParams(req)
	if err != nil {
		var resp Response
		resp.Code = 422
		resp.Msg = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(resp)
		return
	}
	processedRows, err := dataProcessService.ProcessCSV(req)
	if err != nil {
		var resp Response
		resp.Code = 422
		resp.Msg = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		var resp Response
		resp.Code = 200
		resp.Msg = "CSV successfully uploaded. Processed row count:" + strconv.Itoa(processedRows) //add currupted row count too
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}

}

func parseStruct(data []string) *database.PriceData {
	unix, _ := strconv.ParseInt(data[0], 10, 64)
	open, _ := strconv.ParseFloat(data[2], 64)
	hign, _ := strconv.ParseFloat(data[3], 64)
	low, _ := strconv.ParseFloat(data[4], 64)
	close, _ := strconv.ParseFloat(data[5], 64)

	pd := &database.PriceData{
		Unix:       unix,
		Symbol:     data[1],
		OpenPrice:  open,
		HighPrice:  hign,
		LowPrice:   low,
		ClosePrice: close,
	}
	fmt.Println(pd)
	database.Instance.Create(&pd)
	return pd
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

func processParams(req *http.Request) error {
	csvPartFile, header, openErr := req.FormFile("file")
	if openErr != nil {
		return errors.FileOpenCSVError
	}
	if filepath.Ext(header.Filename) != ".csv" {
		return errors.InvalidCSVError
	}

	//close the file at the end
	defer csvPartFile.Close()

	fcsv := csv.NewReader(csvPartFile)
	column, _ := fcsv.Read()

	if len(column) != 6 || consts.CSVHeaderCol1 != column[0] || consts.CSVHeaderCol2 != column[1] || consts.CSVHeaderCol3 != column[2] || consts.CSVHeaderCol4 != column[3] || consts.CSVHeaderCol5 != column[4] || consts.CSVHeaderCol6 != column[5] {
		return errors.InvalidCSVHeaderError
	}
	return nil
}
