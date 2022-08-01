package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string
	ProductID   string
	Description string
	ProductType int
	Total       int
	LeftCount   int
	StockInfo   Stock // 所在仓库
}
