package model

import (
	"time"
)

type Product struct {
	ProductId   int `gorm:"primarykey"`
	ProductName string
	Description string
	ProductType int
	Total       int64
	LeftCount   int64
	Price       int64 // 指导价

	StockID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderProduct struct {
	ProductId   int `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ProductName string
	Description string
	ProductType int
	Total       int64 // 一件商品多份
	Price       int64 // 每个订单的价格允许不一样

	StockID uint   `gorm:"index"`
	OrderID string `gorm:"index"`
}
