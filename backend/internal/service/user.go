package service

import (
	"errors"

	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/pkg/response"
	"github.com/Smash249/xenova/backend/utils"
	"gorm.io/gorm"
)

var UserServicesApp = new(UserService)

type UserService struct {
}

func (UserService) Login(params request.UserLoginReq) (*response.UserLoginRes, error) {
	var user models.User
	err := global.DB.Where("email = ?", params.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	if !utils.CheckPasswordHash(params.Password, user.Password) {
		return nil, errors.New("密码错误")
	}

	accessToken, rereshToken, err := utils.GenerateTokenPair(utils.BaseClaims{
		ID:       int(user.ID),
		UserName: user.UserName,
	})
	if err != nil {
		return nil, errors.New("生成 token 失败")
	}

	result := &response.UserLoginRes{
		User:         user,
		AccessToken:  accessToken,
		ReFreshToken: rereshToken,
	}
	return result, nil
}

func (UserService) Register(params request.UserRegisterReq) error {
	// todo: 校验邮箱
	return global.DB.Create(&models.User{
		UserName: params.UserName,
		Email:    params.Email,
		Password: params.Password,
	}).Error

}
