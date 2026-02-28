package service

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/pkg/response"
	"github.com/Smash249/xenova/backend/utils"
	"gorm.io/gorm"
)

type UserService struct {
}

func (u *UserService) baseLogin(user models.User, params request.UserLoginReq) (*response.UserLoginResp, error) {
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

func (u *UserService) FrontendLogin(params request.UserLoginReq) (*response.UserLoginResp, error) {
	var user models.User
	err := global.DB.Where("email = ?", params.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	if user.IsBanned {
		return nil, errors.New("用户已被封禁")
	}
	return u.baseLogin(user, params)
}

func (u *UserService) AdminLogin(params request.UserLoginReq) (*response.UserLoginResp, error) {
	var user models.User
	err := global.DB.Where("email = ?", params.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	if user.Role != "admin" {
		return nil, errors.New("没有管理员权限")
	}
	return u.baseLogin(user, params)
}

func (u *UserService) Register(params request.UserRegisterReq) error {
	// todo: 校验邮箱
	return global.DB.Create(&models.User{
		UserName: params.UserName,
		Email:    params.Email,
		Password: params.Password,
	}).Error
}

func (u *UserService) GetUserInfo(userID uint) (*models.User, error) {
	var user models.User
	err := global.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserService) UpdateUserInfo(userID uint, params request.UpdateUserInfoReq) (*models.User, error) {
	var user models.User
	err := global.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	user.UserName = params.UserName
	user.Phone = params.Phone
	if err := global.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) AdminUpdateUserInfo(userID uint, params request.UpdateAdminInfoReq) (*models.User, error) {
	var user models.User
	err := global.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	user.UserName = params.UserName
	if err := global.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) ChangePassword(userID uint, params request.ChangePasswordReq) error {
	var user models.User
	err := global.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}
	if !utils.CheckPasswordHash(params.OldPassword, user.Password) {
		return errors.New("旧密码错误")
	}
	hashed, err := utils.GenerateHashPassword(params.NewPassword)
	if err != nil {
		return err
	}
	user.Password = hashed
	return global.DB.Save(&user).Error
}

func (u *UserService) Upload(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", err
	}
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d_%d%s", time.Now().UnixNano(), rand.Intn(10000), ext)
	dst := filepath.Join(uploadDir, filename)
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()
	if _, err := io.Copy(out, src); err != nil {
		return "", err
	}
	fileURL := fmt.Sprintf("/static/%s", filename)
	return fileURL, nil
}
