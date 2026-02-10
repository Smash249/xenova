package models

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/utils"
	"gorm.io/gorm"
)

type User struct {
	global.BaseModel
	UserName string `gorm:"type:varchar(100);uniqueIndex;not null" json:"userName"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}

func (User) TableName() string {
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
