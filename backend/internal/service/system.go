package service

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
)

type SystemService struct {
}

func (SystemService) GetSystemUserList(params request.GetSystemUserListReq) (*global.PaginatorResp[models.User], error) {
	query := global.DB.Model(&models.User{}).Where("role = ?", "user")
	if params.Email != "" {
		query = query.Where("email LIKE ?", "%"+params.Email+"%")
	}
	if params.UserName != "" {
		query = query.Where("user_name LIKE ?", "%"+params.UserName+"%")
	}
	result, err := utils.Paginator[models.User](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (SystemService) BanedUserById(userId uint) error {
	return global.DB.Model(&models.User{}).Where("id = ?", userId).Update("is_banned", true).Error
}
