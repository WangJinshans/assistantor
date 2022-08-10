package model

import (
	"time"
)

// 支付信息
type Payment struct {
	PaymentID string `gorm:"primarykey"`
	Status    string
	Price     int64
	UserId    string
	To        string
	Way       int // 支付方式
	OrderId   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
