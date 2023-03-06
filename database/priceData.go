package database

import (
	"fmt"
	"log"
	"strings"
)

type PriceData struct {
	Unix       int64   `json:"unix"`
	Symbol     string  `json:"symbol" gorm:"size:191"`
	OpenPrice  float64 `json:"open"`
	HighPrice  float64 `json:"high"`
	LowPrice   float64 `json:"low"`
	ClosePrice float64 `json:"close"`
}

func (pd *PriceData) TableName() string {
	return "price_data"
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

func GetAllDetails(pd *PriceData, filter *Filter) []PriceData {
	fmt.Println(filter)
	var res []PriceData
	fields, values := generateQuery(*filter)
	err := Instance.Find(&pd).Select("symbol,unix,open_price,close_price,high_price,low_price").Where(strings.Join(fields, " AND "), values...).Offset((filter.Page - 1) * filter.Count).Limit(filter.Count).Order(filter.SortBy + " " + filter.Order).Scan(&res).Error

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	return res
}

func getSign(signStr string) string {
	switch {
	case signStr == "eq":
		return "="
	case signStr == "lt":
		return "<"
	case signStr == "gt":
		return " > "
	case signStr == "lteq":
		return "<="
	case signStr == "gteq":
		return ">="
	case signStr == "lt":
		return "<"
	default:
		return "="
	}
}

func generateQuery(filter Filter) ([]string, []interface{}) {

	fields := []string{}
	values := []interface{}{}

	if filter.Symbol != "" {
		if strings.Contains(filter.Symbol, ":") {
			splits := strings.Split(filter.Symbol, ":")
			fields = append(fields, "symbol "+getSign(splits[0])+" ?")
			values = append(values, splits[1])
		} else {
			fields = append(fields, "symbol = ?")
			values = append(values, filter.Symbol)
		}
	}

	if filter.HighPrice != "" {
		if strings.Contains(filter.HighPrice, ":") {
			splits := strings.Split(filter.HighPrice, ":")
			fields = append(fields, "high_price "+getSign(splits[0])+" ?")
			values = append(values, splits[1])
		} else {
			fields = append(fields, "high_price = ?")
			values = append(values, filter.HighPrice)
		}
	}

	if filter.LowPrice != "" {
		if strings.Contains(filter.LowPrice, ":") {
			splits := strings.Split(filter.LowPrice, ":")
			fields = append(fields, "low_price "+getSign(splits[0])+" ?")
			values = append(values, splits[1])
		} else {
			fields = append(fields, "low_price = ?")
			values = append(values, filter.LowPrice)
		}
	}

	if filter.OpenPrice != "" {
		if strings.Contains(filter.OpenPrice, ":") {
			splits := strings.Split(filter.OpenPrice, ":")
			fields = append(fields, "open_price "+getSign(splits[0])+" ?")
			values = append(values, splits[1])
		} else {
			fields = append(fields, "open_price = ?")
			values = append(values, filter.OpenPrice)
		}
	}

	if filter.ClosePrice != "" {
		if strings.Contains(filter.ClosePrice, ":") {
			splits := strings.Split(filter.ClosePrice, ":")
			fields = append(fields, "close_price "+getSign(splits[0])+" ?")
			values = append(values, splits[1])
		} else {
			fields = append(fields, "close_price = ?")
			values = append(values, filter.ClosePrice)
		}
	}

	return fields, values
}
