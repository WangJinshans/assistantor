package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string
	Description string
	ProductType int
	Total       int
	LeftCount   int

	StockID uint
}

type OrderProduct struct {
	gorm.Model
	ProductName string
	Description string
	ProductType int
	Total       int
	LeftCount   int

	StockID     uint   `gorm:"index"`
	OrderID     string `gorm:"index"`
}
