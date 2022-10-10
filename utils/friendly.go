package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
)

func ShowNumberFriendly(data float64) (result float64, u string) {
	value := math.Abs(data)
	if value > 100000000 {
		result, _ = decimal.NewFromFloat(data / 100000000).Round(2).Float64()
		u = "亿"
	} else if value > 10000 {
		result, _ = decimal.NewFromFloat(data / 10000).Round(2).Float64()
		u = "万"
	}
	return
}

func ShowNumberFriendlyString(data float64) (result string) {
	value, u := ShowNumberFriendly(data)
	result = fmt.Sprintf("%.2f%s", value, u)
	return
}

func ShowNumberFriendlyNumberWithUnit(data float64, uit int) (result float64) {
	if uit == 1 {
		result, _ = decimal.NewFromFloat(data / 100000000).Round(2).Float64()
	} else {
		result, _ = decimal.NewFromFloat(data / 10000).Round(2).Float64()
	}
	return
}
