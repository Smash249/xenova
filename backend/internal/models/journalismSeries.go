package models

import "github.com/Smash249/xenova/backend/internal/global"

type JournalismSeries struct {
	global.BaseModel
	Name        string       `gorm:"type:varchar(100);uniqueIndex;not null;comment:新闻系列名称" json:"name"`
	Description string       `gorm:"type:text;comment:新闻系列描述" json:"description"`
	Journalisms []Journalism `gorm:"foreignKey:SeriesID" json:"-"`
}

func (JournalismSeries) TableName() string {
	return "journalism_series"
}
