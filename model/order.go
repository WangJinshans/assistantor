package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID     string `gorm:"primarykey"`
	OrderStatus int

	UserID      string
	AddressID   uint
}
