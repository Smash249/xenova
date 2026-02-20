package service

import (
	"errors"
	"fmt"

	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/pkg/response"
	"github.com/Smash249/xenova/backend/utils"
	"gorm.io/gorm"
)

type UserService struct {
}

func (UserService) Login(params request.UserLoginReq) (*response.UserLoginResp, error) {
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

	accessToken, refreshToken, err := utils.GenerateTokenPair(utils.BaseClaims{
		ID:       user.ID,
		UserName: user.UserName,
		Role:     user.Role,
	})
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("生成 token 失败")
	}

	result := &response.UserLoginResp{
		User:         user,
		AccessToken:  accessToken,
		ReFreshToken: refreshToken,
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
