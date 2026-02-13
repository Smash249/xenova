package models

import (
	"github.com/Smash249/xenova/backend/internal/global"
)

type Journalism struct {
	global.BaseModel
	Title       string     `gorm:"type:varchar(200);not null;comment:新闻标题" json:"title"`
	Content     string     `gorm:"type:text;comment:新闻内容" json:"content"`
	Cover       string     `gorm:"type:varchar(100);not null;comment:新闻封面图" json:"cover"`
	Author      string     `gorm:"type:varchar(100);comment:作者" json:"author"`
	Source      string     `gorm:"type:varchar(100);comment:来源" json:"source"`
	Tags        StringList `gorm:"type:text;comment:标签" json:"tags"`
	SeriesID    uint       `gorm:"comment:新闻系列Id" json:"seriesId"`
}

func (Journalism) TableName() string {
	return "journalism"
}
