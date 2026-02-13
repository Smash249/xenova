package models

import "github.com/Smash249/xenova/backend/internal/global"

type Software struct {
	global.BaseModel
	Name        string `gorm:"not null;comment:软件名称" json:"name"`
	Description string `gorm:"type:text;comment:软件描述" json:"description"`
	Version     string `gorm:"not null;comment:版本号" json:"version"`
	FileSize    string `gorm:"not null;comment:文件大小" json:"file_size"`
	FileURL     string `gorm:"not null;comment:文件下载地址" json:"file_url"`
	IsHot       bool   `gorm:"default:false;comment:是否热门" json:"is_hot"`
	CategoryID  uint   `gorm:"not null;comment:分类ID" json:"category_id"`
}

func (Software) TableName() string {
	return "software"
}
