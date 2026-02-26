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

func (SystemService) GetSystemMessageList(params request.GetSystemMessageListReq) (*global.PaginatorResp[models.Message], error) {
	query := global.DB.Model(&models.Message{})
	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if params.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+params.Phone+"%")
	}
	result, err := utils.Paginator[models.Message](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (SystemService) CreateSystemMessage(userId uint, params request.CreateSystemMessageReq) error {
	global.DB.Create(models.Message{
		UserID:  userId,
		Content: params.Content,
		Company: params.Company,
		Email:   params.Email,
		Phone:   params.Phone,
	})
	return nil
}

func (SystemService) UpdateSystemMessage(params request.UpdateSystemMessageReq) error {
	return global.DB.Model(&models.Message{}).Where("id = ?", params.ID).Updates(models.Message{
		Content: params.Content,
		Company: params.Company,
		Email:   params.Email,
		Phone:   params.Phone,
	}).Error
}

func (SystemService) DeleteSystemMessage(params request.DeleteSystemMessageReq) error {
	return global.DB.Delete(&models.Message{}, params.IdList).Error
}

func (SystemService) GetSystemJobPositionList(params request.GetSystemJobPositionListReq) (*global.PaginatorResp[models.JobPosition], error) {
	query := global.DB.Model(&models.JobPosition{})
	if params.Title != "" {
		query = query.Where("title LIKE ?", "%"+params.Title+"%")
	}
	if params.Department != "" {
		query = query.Where("department LIKE ?", "%"+params.Department+"%")
	}
	if params.Status != 0 {
		query = query.Where("status = ?", params.Status)
	}
	result, err := utils.Paginator[models.JobPosition](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (SystemService) CreateSystemJobPosition(params request.CreateSystemJobPositionReq) error {
	return global.DB.Model(&models.JobPosition{}).Create(models.JobPosition{
		Title:            params.Title,
		Department:       params.Department,
		Location:         params.Location,
		Headcount:        params.Headcount,
		Experience:       params.Experience,
		Education:        params.Education,
		SalaryRange:      params.SalaryRange,
		Responsibilities: params.Responsibilities,
		Requirements:     params.Requirements,
	}).Error
}

func (SystemService) UpdateSystemJobPosition(params request.UpdateSystemJobPositionReq) error {
	return global.DB.Model(&models.JobPosition{}).Where("id = ?", params.ID).Updates(models.JobPosition{
		Title:            params.Title,
		Department:       params.Department,
		Location:         params.Location,
		Headcount:        params.Headcount,
		Experience:       params.Experience,
		Education:        params.Education,
		SalaryRange:      params.SalaryRange,
		Responsibilities: params.Responsibilities,
		Requirements:     params.Requirements,
		Status:           params.Status,
		Sort:             params.Sort,
	}).Error
}

func (SystemService) DeleteSystemJobPosition(params request.DeleteSystemJobPositionReq) error {
	return global.DB.Delete(&models.JobPosition{}, params.IdList).Error
}
