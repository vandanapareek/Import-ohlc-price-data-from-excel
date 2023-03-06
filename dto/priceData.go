package dto

type CsvHeaderPriceData struct {
	Unix   int64   `json:"UNIX"`
	Symbol string  `json:"SYMBOL"`
	Open   float64 `json:"OPEN"`
	High   float64 `json:"HIGH"`
	Low    float64 `json:"LOW"`
	Close  float64 `json:"CLOSE"`
}
