package model

import "time"

// 购物车
type CartProduct struct {
	ID           uint `gorm:"primarykey"`
	UserId       string
	StoreId      string
	StoreName    string // 店名
	ProductId    string // StoreProduct
	ProductUrl   string // 链接
	ProductImage string // 封面
	Price        int64  // 显示价格
	Count        int64  // 数量
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
