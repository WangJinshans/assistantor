package model

import (
	"time"
)

// 评价
type ProductComment struct {
	ID             uint `gorm:"primarykey"`
	ProductId      string
	CommentId      int // 消息ID
	UserId         string
	Comment        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	AttachmentInfo []MediaResource
}
