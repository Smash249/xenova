package models

import "github.com/Smash249/xenova/backend/internal/global"

type ProductSeries struct {
	global.BaseModel
	Name        string    `gorm:"type:varchar(100);uniqueIndex;not null;comment:产品系列名称" json:"name"`
	Description string    `gorm:"type:text;comment:产品系列描述" json:"description"`
	Products    []Product `gorm:"foreignKey:SeriesID" json:"-"`
}

func (ProductSeries) TableName() string {
	return "product_series"
}
