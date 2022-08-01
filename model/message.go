package model

import "gorm.io/gorm"

// default user 机器人
type Message struct {
	gorm.Model
	FromUserID  string
	ToUserID    string
	Content     string
	MessageType int
	Attachment  []MediaResource // 资源文件
}

type MediaResource struct {
	gorm.Model
	ResourceID   string
	ResourceName string
	ResourceType int // url 静态
	ResourceUrl  string
	ResourcePath string
}
