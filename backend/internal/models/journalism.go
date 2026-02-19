package models

import "github.com/Smash249/xenova/backend/internal/global"

type Journalism struct {
	global.BaseModel
	Title     string `gorm:"not null;comment:新闻标题" json:"title"`
	Content   string `gorm:"type:text;comment:新闻内容" json:"content"`
	SeriesID  uint   `gorm:"comment:新闻系列 id" json:"series_id"`
	ViewCount uint   `gorm:"default:0;comment:浏览量" json:"view_count"`
}

func (Journalism) TableName() string {
	return "journalism"
}
