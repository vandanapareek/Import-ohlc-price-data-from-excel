package errors

import (
	"fmt"
)

type CSVError struct {
	Code          int
	Message       string
	ResultCode    string
	ResultStatus  string
	ResultMessage string
}

// CSVError
func (err CSVError) Error() string {
	return fmt.Sprintf("error_code = %v, error_message = %v", err.Code, err.Message)
}

func CreateError(code int, message string) CSVError {
	return CSVError{Code: code, Message: message}
}

var Success = CSVError{Code: 200, Message: "success"}
var InvalidCSVError = CSVError{Code: 401, Message: "The CSV file is invalid."}
var FileOpenCSVError = CSVError{Code: 401, Message: "Unable to open csv file."}
var InvalidCSVHeaderError = CSVError{Code: 401, Message: "Invalid file headers, expecting as [UNIX,SYMBOL,OPEN,HIGH,LOW,CLOSE]."}
var DatabaseConnectionError = CSVError{Code: 502, Message: "Error establishing a database connection "}
var InvalidParamError = CSVError{Code: 401, Message: "Invalid Params"}
var NoRecordFoundError = CSVError{Code: 401, Message: "No record found"}
