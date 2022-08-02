package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID     string `gorm:"primarykey;index;size:32"`
	OrderStatus int
	UserID      string `gorm:"index"`
	AddressID   uint

	//UserInfo    User            `gorm:"foreignKey:UserID;references:UserId"`
	//AddressInfo DeliveryAddress `gorm:"foreignKey:AddressID;references:ID"`
	//ProductList []OrderProduct  `gorm:"foreignKey:OrderID;references:OrderID"`
}
