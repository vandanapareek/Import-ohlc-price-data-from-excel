package csvdata

import (
	"encoding/csv"
	"fmt"
	"go-price-data/database"
	"go-price-data/dto"
	"go-price-data/errors"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var err error

type ICsvService interface {
	ProcessCSV(req *http.Request) (int, error)
	SearchDetails(filterStruct dto.Filter) ([]database.PriceData, error)
}

type CsvStruct struct {
}

func (csvStruct CsvStruct) SearchDetails(filterStruct dto.Filter) ([]database.PriceData, error) {
	res, err := database.SearchDetails(filterStruct)
	fmt.Printf("hahahhahahahaha %d", len(res))
	if err != nil {
		return res, err
	}
	return res, nil
}

//ProcessCSV reads async rows from a CSV file by initializing fixed number of workers and
//multiple goroutines, each of workers handling a chunk of data from the file.
func (csvStruct CsvStruct) ProcessCSV(req *http.Request) (int, error) {

	//init val
	rs := make([]*database.PriceData, 0)
	numWps := 100
	jobs := make(chan []string, numWps)
	res := make(chan *database.PriceData)

	csvPartFile, _, _ := req.FormFile("file")

	//close the file at the end
	defer csvPartFile.Close()

	fcsv := csv.NewReader(csvPartFile)
	fcsv.Read()

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- *database.PriceData) {
		for {
			select {
			case job, ok := <-jobs: // you must check for readable state of the channel.

				if !ok {
					return
				}
				results <- parseStruct(job)
			}
		}
	}

	if err != nil {
		fmt.Print(err)
	}
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}

	// init workers
	for w := 0; w < numWps; w++ {

		wg.Add(1)
		go func() {
			// this line will exec when chan `res` processed output at line 94 (func worker: line 47)
			defer wg.Done()
			worker(jobs, res)
		}()

	}

	go func() {
		lineNum := 0
		for {
			lineNum++
			column, err := fcsv.Read()

			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				break
			}
			jobs <- column
		}
		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	go func() {
		wg.Wait()
		close(res) // when you close(res) it breaks the below loop.
	}()

	for r := range res {
		rs = append(rs, r)
	}

	fmt.Println("Count Concu ", len(rs))

	if len(rs) == 0 {
		return 0, errors.InvalidCSVError
	}
	return len(rs), nil
}

//parse content of each row and save to database
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
