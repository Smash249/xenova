package models

import (
	"time"

	"github.com/Smash249/xenova/backend/internal/global"
)

type CompanyHonor struct {
	global.BaseModel
	Title       string    `gorm:"type:varchar(255);not null;comment:'荣誉名称'"`
	IssueDate   time.Time `gorm:"type:date;comment:'颁发日期'"`
	CertNo      string    `gorm:"type:varchar(100);comment:'证书编号'"`
	Issuer      string    `gorm:"type:varchar(255);comment:'颁发机构'"`
	Image       string    `gorm:"type:varchar(500);comment:'图片链接'"`
	Description string    `gorm:"type:text;comment:'荣誉介绍'"`
}

func (CompanyHonor) TableName() string {
	return "company_honor"
}

type LoveActivity struct {
	global.BaseModel
	Title        string    `gorm:"type:varchar(255);not null;comment:'活动标题'"`
	Location     string    `gorm:"type:varchar(255);comment:'活动地点'"`
	Participants int       `gorm:"type:int;default:0;comment:'参与人数'"`
	Cover        string    `gorm:"type:varchar(500);comment:'封面图链接'"`
	Summary      string    `gorm:"type:varchar(500);comment:'活动摘要'"`
	Content      string    `gorm:"type:longtext;comment:'活动纪实(长文本)'"`
	ActivityDate time.Time `gorm:"type:date;comment:'活动日期'"`
}

func (LoveActivity) TableName() string {
	return "love_activity"
}

type CompanyPatent struct {
	global.BaseModel
	Title    string    `gorm:"type:varchar(255);not null;comment:'专利名称'"`
	Type     string    `gorm:"type:varchar(50);comment:'专利类型(如:发明专利/实用新型/外观设计)'"`
	PatentNo string    `gorm:"type:varchar(100);comment:'专利号/申请号'"`
	AuthDate time.Time `gorm:"type:date;comment:'授权公告日'"`
	Inventor string    `gorm:"type:varchar(255);comment:'发明人/设计人'"`
	Image    string    `gorm:"type:varchar(500);comment:'专利图片链接'"`
	Summary  string    `gorm:"type:text;comment:'摘要说明'"`
}

func (CompanyPatent) TableName() string {
	return "company_patent"
}
