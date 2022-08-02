package model

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	StoreType int
	StoreName string
	ScaleMax  int
	ScaleMin  int

	AddressID uint   `gorm:"index"`
	UserID    string `gorm:"index"`

	//// 外键
	//StoreAddress Address `gorm:"foreignKey:AddressID;references:ID"`  // 地址
	//Owner        User    `gorm:"foreignKey:UserId;references:UserID"` // 店主
}
