package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string
	Description string
	ProductType int
	Total       int
	LeftCount   int

	// 外键
	StockID   uint  `gorm:"index"`
	//StockInfo Stock `gorm:"foreignKey:StockID;references:ID"` // 所在仓库

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

	// 外键
	//StockInfo Stock `gorm:"foreignKey:StockID;references:ID"` // 所在仓库

}
