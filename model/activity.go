package model

import "gorm.io/gorm"

// 活动

type StoreActivity struct {
	gorm.Model
	Description string
	Discount    int64
}

// 优惠券
type Voucher struct {
	gorm.Model
	Description    string
	Discount       int64
	RequireConsume int64
}
