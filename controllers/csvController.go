package controllers

import (
	"encoding/csv"
	"encoding/json"
	"go-price-data/consts"
	"go-price-data/errors"
	csvservice "go-price-data/services/csvService"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	csvService csvservice.ICsvService = csvservice.ICsvService(csvservice.CsvStruct{})
)

func ReadCsv(w http.ResponseWriter, req *http.Request) {

	//validate csv
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

	//process CSV data
	processedRows, err := csvService.ProcessCSV(req)

	if err != nil {
		//CSV have some errors
		var resp Response
		resp.Code = 422
		resp.Msg = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		//CSV processed successfully
		var resp Response
		resp.Code = 200
		resp.Msg = "CSV successfully uploaded. Processed row count:" + strconv.Itoa(processedRows)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}

}

//validate CSV file and its header
func processParams(req *http.Request) error {
	csvPartFile, header, openErr := req.FormFile("file")
	if openErr != nil {
		return errors.FileOpenCSVError
	}
	if filepath.Ext(header.Filename) != ".csv" {
		return errors.InvalidCSVError
	}

	//close the file
	defer csvPartFile.Close()

	//validate csv header
	fcsv := csv.NewReader(csvPartFile)
	column, _ := fcsv.Read()

	if len(column) != 6 || consts.CSVHeaderCol1 != column[0] || consts.CSVHeaderCol2 != column[1] || consts.CSVHeaderCol3 != column[2] || consts.CSVHeaderCol4 != column[3] || consts.CSVHeaderCol5 != column[4] || consts.CSVHeaderCol6 != column[5] {
		return errors.InvalidCSVHeaderError
	}
	return nil
}
