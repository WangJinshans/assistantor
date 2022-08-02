package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Lng          float64
	Lat          float64
	ProvinceName string
	DistrictName string
	StreetName   string
	ExtraInfo    string
}

type DeliveryAddress struct {
	gorm.Model
	Lng          float64
	Lat          float64
	AddressType  int // 是否默认
	ProvinceName string
	DistrictName string
	StreetName   string
	ExtraInfo    string

	UserID string // 外键
}
