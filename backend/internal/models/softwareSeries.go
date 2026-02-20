package models

import (
	"github.com/Smash249/xenova/backend/internal/global"
)

type SoftwareSeries struct {
	global.BaseModel
	Name        string     `gorm:"unique;not null;comment:分类名称" json:"name"`
	Description string     `gorm:"type:text;comment:分类描述" json:"description"`
	Software    []Software `gorm:"foreignKey:SeriesID" json:"-"`
}

func (SoftwareSeries) TableName() string {
	return "software_series"
}
