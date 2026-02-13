package models

import "github.com/Smash249/xenova/backend/internal/global"

type SoftwareCategory struct {
	global.BaseModel
	Name     string     `gorm:"unique;not null;comment:分类名称" json:"name"`
	Sort     uint       `gorm:"default:0;comment:排序权重" json:"sort"`
	Software []Software `gorm:"foreignKey:CategoryID" json:"-"`
}

func (SoftwareCategory) TableName() string {
	return "software_categories"
}
