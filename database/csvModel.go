package database

import (
	"fmt"
	"go-price-data/dto"
	"go-price-data/errors"
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

func SearchDetails(filter dto.Filter) ([]PriceData, error) {
	var pd *PriceData
	var res []PriceData
	fields, values := generateQuery(filter)
	err := Instance.Find(&pd).Select("symbol,unix,open_price,close_price,high_price,low_price").Where(strings.Join(fields, " AND "), values...).Offset((filter.Page - 1) * filter.Count).Limit(filter.Count).Order(filter.SortBy + " " + filter.Order).Scan(&res).Error

	if err != nil {
		fmt.Println("I am failed")

		return res, errors.DatabaseConnectionError
	}
	fmt.Println("I am successfull")
	return res, nil
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

func generateQuery(filter dto.Filter) ([]string, []interface{}) {

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
