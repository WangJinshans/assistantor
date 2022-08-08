package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID string `gorm:"primarykey"`

	OrderType   int `json:"order_type"` // 订单类型  打赏/会员/商品交易/酒店/出行/吃住
	OrderStatus string
	Price       int64 // 分

	UserID    string
	AddressID uint
}

type OrderRequest struct {
	UserId      string `json:"user_id"`
	OrderId     string `json:"order_id"`
	ProductList []struct {
		ProductId   string `json:"product_id"`
		ProductName string `json:"product_name"`
		Count       int64  `json:"count"`
	} `json:"product_list"`
}
