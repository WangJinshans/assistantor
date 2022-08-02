package model

import "gorm.io/gorm"

// 库存
type Stock struct {
	gorm.Model
	Address    string
	StockName  string
}
