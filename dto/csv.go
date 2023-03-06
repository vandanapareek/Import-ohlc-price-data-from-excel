package dto

type CsvHeaderPriceData struct {
	Unix   int64   `json:"UNIX"`
	Symbol string  `json:"SYMBOL"`
	Open   float64 `json:"OPEN"`
	High   float64 `json:"HIGH"`
	Low    float64 `json:"LOW"`
	Close  float64 `json:"CLOSE"`
}

type Filter struct {
	Page   int    `schema:"page"`
	Count  int    `schema:"count"`
	SortBy string `schema:"sortby"`
	Order  string `schema:"order"`
	//Price specific filters
	Symbol     string `schema:"symbol"`
	HighPrice  string `schema:"high_price"`
	LowPrice   string `schema:"low_price"`
	OpenPrice  string `schema:"open_price"`
	ClosePrice string `schema:"close_price"`
}

func SetDefault() Filter {
	return Filter{
		Page:   1,
		Count:  10,
		SortBy: "symbol",
		Order:  "asc",
	}
}
