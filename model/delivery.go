package model

import "gorm.io/gorm"

type DeliveryInfo struct {
	gorm.Model
	OrderId           string
	UserId            string
	DeliveryType      int
	DeliveryStatus    int
	DeliveryAddressId int
	RiderId           string // 骑手id
}
