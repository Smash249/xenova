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
	result, err := utils.Paginator[models.User](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}
