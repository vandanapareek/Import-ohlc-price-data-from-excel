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
