package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID       uint
	ProvinceName string
	DistrictName string
	StreetName   string
	ExtraInfo    string
}
