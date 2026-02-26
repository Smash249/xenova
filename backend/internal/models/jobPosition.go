package models

import "github.com/Smash249/xenova/backend/internal/global"

// JobPosition 招聘职位模型
type JobPosition struct {
	global.BaseModel
	Title            string `gorm:"type:varchar(100);not null;comment:职位名称" json:"title"`
	Department       string `gorm:"type:varchar(50);default:'';comment:所属部门" json:"department"`
	Location         string `gorm:"type:varchar(50);not null;default:'厦门';comment:工作地点" json:"location"`
	Headcount        string `gorm:"type:varchar(20);not null;default:'若干';comment:招聘人数" json:"headcount"`
	Experience       string `gorm:"type:varchar(50);default:'不限';comment:工作经验要求(如: 1-3年)" json:"experience"`
	Education        string `gorm:"type:varchar(50);default:'不限';comment:学历要求(如: 本科及以上)" json:"education"`
	SalaryRange      string `gorm:"type:varchar(50);default:'面议';comment:薪资范围(如: 10k-15k)" json:"salaryRange"`
	Responsibilities string `gorm:"type:text;not null;comment:岗位职责(支持富文本或换行符)" json:"responsibilities"`
	Requirements     string `gorm:"type:text;not null;comment:任职要求(支持富文本或换行符)" json:"requirements"`
	Status           int    `gorm:"type:tinyint(1);default:1;comment:状态(1:招聘中, 0:已停止/下线)" json:"status"`
	Sort             int    `gorm:"type:int;default:0;comment:排序字段(数值越小越靠前)" json:"sort"`
}

func (JobPosition) TableName() string {
	return "sys_job_positions"
}
