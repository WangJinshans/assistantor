package model

import (
	"time"
)


type StoreProduct struct {
	ProductId    string `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProductName  string
	Description  string
	ProductType  int
	LockCount    int64 // 临时锁定(未支付)
	Count        int64 //
	LeftCount    int64 // 剩余
	Price        int64 // 每个订单的价格允许不一样
	DefaultPrice int64 // 指导价

	StoreID string
	OrderID string
}

// 商品订单
type OrderProduct struct {
	ProductId   string `gorm:"primarykey"`
	ProductName string

	CreatedAt time.Time
	UpdatedAt time.Time

	Count int64 // 一件商品多份
	Price int64 // 每个订单的价格允许不一样

	StoreId      string
	OrderId      string
	UserId       string // 付款方
	AcceptUserId string // 收货方
}
