package models

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/utils"
	"gorm.io/gorm"
)

type User struct {
	global.BaseModel
	UserName string `gorm:"type:varchar(100);uniqueIndex;not null;comment:用户名" json:"userName"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null;comment:邮箱" json:"email"`
	Password string `gorm:"type:varchar(255);not null;comment:密码" json:"-"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashed, err := utils.GenerateHashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}
