package models

import (
	"github.com/Smash249/xenova/backend/internal/global"
)

type Product struct {
	global.BaseModel
	Name        string     `gorm:"type:varchar(100);not null;comment:产品名称" json:"name"`
	Description string     `gorm:"type:text;comment:产品描述" json:"description"`
	Cover       string     `gorm:"type:varchar(100);not null;comment:产品图片" json:"cover"`
	Previews    StringList `gorm:"type:text;comment:产品预览图" json:"previews"`
	IsHot       bool       `gorm:"default:false;comment:是否热门" json:"is_hot"`
	SeriesID    uint       `gorm:"comment:产品系列Id" json:"series_id"`
}

func (Product) TableName() string {
	return "product"
}
