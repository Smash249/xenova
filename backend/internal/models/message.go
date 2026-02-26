package models

import "github.com/Smash249/xenova/backend/internal/global"

type Message struct {
	global.BaseModel
	UserID  uint   `gorm:"not null;comment:用户ID" json:"userId"`
	Content string `gorm:"type:text;not null;comment:消息内容" json:"content"`
	Company string `gorm:"type:varchar(100);comment:公司名称" json:"company"`
	Email   string `gorm:"type:varchar(100);comment:邮箱" json:"email"`
	Phone   string `gorm:"type:varchar(20);comment:手机号" json:"phone"`
}

func (Message) TableName() string {
	return "message"
}
