package model

import "gorm.io/gorm"

type DeliveryInfo struct {
	gorm.Model
	OrderId           string
	DeliveryType      int
	DeliveryStatus    int
	DeliveryAddressId int
}
