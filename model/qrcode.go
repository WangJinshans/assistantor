package model

import "gorm.io/gorm"

type QrCodeInfo struct {
	gorm.Model
	Id        int64
	QrCodeId  string
	IsValid   string
	ImagePath string
	Status    int
	UserId    string // 用户Id
	TimeStamp int64  // 创建时间
	Content   string // base64 字符串
}
