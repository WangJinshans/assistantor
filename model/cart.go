package model

import "time"

// 购物车
type Cart struct {
	ID        uint `gorm:"primarykey"`
	UserId    string
	StoreId   string
	ProductId string // StoreProduct
	Count     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
