package models

import "github.com/Smash249/xenova/backend/internal/global"

type AccessorySeries struct {
	global.BaseModel
	Name        string      `gorm:"type:varchar(100);uniqueIndex;not null;comment:配件系列名称" json:"name"`
	Description string      `gorm:"type:text;comment:配件系列描述" json:"description"`
	Accessories []Accessory `gorm:"foreignKey:SeriesID" json:"-"`
}

func (AccessorySeries) TableName() string {
	return "accessory_series"
}
