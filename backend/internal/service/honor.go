package service

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
)

// ==========================================
// 公司荣誉 (CompanyHonor) 及其 Service
// ==========================================

type HonorService struct{}

// GetCompanyHonorList 分页获取公司荣誉列表
func (h *HonorService) GetCompanyHonorList(params request.GetCompanyHonorListReq) (*global.PaginatorResp[models.CompanyHonor], error) {
	query := global.DB.Model(&models.CompanyHonor{})
	if params.Title != "" {
		query = query.Where("title LIKE ?", "%"+params.Title+"%")
	}
	result, err := utils.Paginator[models.CompanyHonor](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCompanyHonor 创建公司荣誉
func (h *HonorService) CreateCompanyHonor(params request.CreateCompanyHonorReq) error {
	return global.DB.Create(&models.CompanyHonor{
		Title:       params.Title,
		IssueDate:   params.IssueDate,
		CertNo:      params.CertNo,
		Issuer:      params.Issuer,
		Image:       params.Image,
		Description: params.Description,
	}).Error
}

// UpdateCompanyHonor 更新公司荣誉
func (h *HonorService) UpdateCompanyHonor(params request.UpdateCompanyHonorReq) error {
	return global.DB.Where("id = ?", params.ID).Updates(&models.CompanyHonor{
		Title:       params.Title,
		IssueDate:   params.IssueDate,
		CertNo:      params.CertNo,
		Issuer:      params.Issuer,
		Image:       params.Image,
		Description: params.Description,
	}).Error
}

// DeleteCompanyHonor 删除公司荣誉
func (h *HonorService) DeleteCompanyHonor(params request.DeleteCompanyHonorReq) error {
	return global.DB.Delete(&models.CompanyHonor{}, params.IdList).Error
}

// ==========================================
// 爱心活动 (LoveActivity) 及其 Service
// ==========================================

// GetActivityList 分页获取爱心活动列表
func (h *HonorService) GetActivityList(params request.GetLoveActivityListReq) (*global.PaginatorResp[models.LoveActivity], error) {
	query := global.DB.Model(&models.LoveActivity{})
	result, err := utils.Paginator[models.LoveActivity](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateActivity 创建爱心活动
func (h *HonorService) CreateActivity(params request.CreateLoveActivityReq) error {
	return global.DB.Create(&models.LoveActivity{
		Title:        params.Title,
		Location:     params.Location,
		Participants: params.Participants,
		Cover:        params.Cover,
		Summary:      params.Summary,
		Content:      params.Content,
		ActivityDate: params.ActivityDate,
	}).Error
}

// UpdateActivity 更新爱心活动
func (h *HonorService) UpdateActivity(params request.UpdateLoveActivityReq) error {
	return global.DB.Where("id = ?", params.ID).Updates(&models.LoveActivity{
		Title:        params.Title,
		Location:     params.Location,
		Participants: params.Participants,
		Cover:        params.Cover,
		Summary:      params.Summary,
		Content:      params.Content,
		ActivityDate: params.ActivityDate,
	}).Error
}

// DeleteActivity 删除爱心活动
func (h *HonorService) DeleteActivity(params request.DeleteLoveActivityReq) error {
	return global.DB.Delete(&models.LoveActivity{}, params.IdList).Error
}

// ==========================================
// 公司专利 (CompanyPatent) 及其 Service
// ==========================================

// GetPatentList 分页获取公司专利列表
func (h *HonorService) GetPatentList(params request.GetCompanyPatentListReq) (*global.PaginatorResp[models.CompanyPatent], error) {
	query := global.DB.Model(&models.CompanyPatent{})
	result, err := utils.Paginator[models.CompanyPatent](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreatePatent 创建公司专利
func (h *HonorService) CreatePatent(params request.CreateCompanyPatentReq) error {
	return global.DB.Create(&models.CompanyPatent{
		Title:    params.Title,
		Type:     params.Type,
		PatentNo: params.PatentNo,
		AuthDate: params.AuthDate,
		Inventor: params.Inventor,
		Image:    params.Image,
		Summary:  params.Summary,
	}).Error
}

// UpdatePatent 更新公司专利
func (h *HonorService) UpdatePatent(params request.UpdateCompanyPatentReq) error {
	return global.DB.Where("id = ?", params.ID).Updates(&models.CompanyPatent{
		Title:    params.Title,
		Type:     params.Type,
		PatentNo: params.PatentNo,
		AuthDate: params.AuthDate,
		Inventor: params.Inventor,
		Image:    params.Image,
		Summary:  params.Summary,
	}).Error
}

// DeletePatent 删除公司专利
func (h *HonorService) DeletePatent(params request.DeleteCompanyPatentReq) error {
	return global.DB.Delete(&models.CompanyPatent{}, params.IdList).Error
}
