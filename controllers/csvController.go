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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}
	processedRows, err := dataProcessService.ProcessCSV(req)
	if err != nil {
		var resp Response
		resp.Code = 422
		resp.Msg = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		var resp Response
		resp.Code = 200
		resp.Msg = "CSV successfully uploaded. Processed row count:" + strconv.Itoa(processedRows) //add currupted row count too
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}
	// rs := make([]*database.PriceData, 0)
	// numWps := 100
	// jobs := make(chan []string, numWps)
	// res := make(chan *database.PriceData)

	// var wg sync.WaitGroup
	// worker := func(jobs <-chan []string, results chan<- *database.PriceData) {
	// 	for {
	// 		select {
	// 		case job, ok := <-jobs: // you must check for readable state of the channel.

	// 			if !ok {
	// 				return
	// 			}
	// 			results <- parseStruct(job)
	// 			//parseStruct(job)
	// 		}
	// 	}
	// }

	// // init workers
	// //it is runnung 100 times even if small csv file
	// for w := 0; w < numWps; w++ {

	// 	wg.Add(1)
	// 	go func() {
	// 		// this line will exec when chan `res` processed output at line 107 (func worker: line 71)
	// 		defer wg.Done()
	// 		worker(jobs, res)
	// 	}()

	// }

	// go func() {
	// 	lineNum := 0
	// 	for {
	// 		lineNum++
	// 		column, err := fcsv.Read()

	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		if err != nil {
	// 			fmt.Println("ERROR: ", err.Error())
	// 			break
	// 		}
	// 		jobs <- column
	// 	}
	// 	close(jobs) // close jobs to signal workers that no more job are incoming.
	// }()

	// go func() {
	// 	wg.Wait()
	// 	close(res) // when you close(res) it breaks the below loop.
	// }()

	// for r := range res {
	// 	rs = append(rs, r)
	// }
	// fmt.Println("Count Concu ", len(rs))

	// if len(rs) == 0 {
	// 	var resp ResponseFailure
	// 	resp.Code = 422
	// 	resp.Msg = "No data processed. Either the CSV file is invalid or empty!"
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// } else {
	// 	var resp ResponseFailure
	// 	resp.Code = 200
	// 	resp.Msg = "CSV successfully uploaded. Processed row count:" + strconv.Itoa(len(rs)) //add currupted row count
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }

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
		// fmt.Println("File openErr error222", openErr)
		// var resp ResponseFailure
		// resp.Code = 422
		// resp.Msg = "Please upload a valid file"
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(resp)
		// return
	}
	if filepath.Ext(header.Filename) != ".csv" {
		return errors.InvalidCSVError
		// var resp ResponseFailure
		// resp.Code = 422
		// resp.Msg = "Invalid file"
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(resp)
		// return
	}

	//close the file at the end
	defer csvPartFile.Close()

	fcsv := csv.NewReader(csvPartFile)
	column, _ := fcsv.Read()

	if len(column) != 6 || consts.CSVHeaderCol1 != column[0] || consts.CSVHeaderCol2 != column[1] || consts.CSVHeaderCol3 != column[2] || consts.CSVHeaderCol4 != column[3] || consts.CSVHeaderCol5 != column[4] || consts.CSVHeaderCol6 != column[5] {
		return errors.InvalidCSVHeaderError
		// var resp ResponseFailure
		// resp.Code = 422
		// resp.Msg = "Invalid file headers, expecting as [UNIX,SYMBOL,OPEN,HIGH,LOW,CLOSE]"
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(resp)
		// return
	}
	return nil
}
