package service

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
)

type SoftwareService struct{}

// ==================== SoftwareCategory 软件分类 CRUD ====================

// GetSoftwareCategoryList 获取软件分类列表
func (s *SoftwareService) GetSoftwareCategoryList() ([]models.SoftwareCategory, error) {
	var categoryList []models.SoftwareCategory
	err := global.DB.Order("sort DESC").Find(&categoryList).Error
	return categoryList, err
}

// CreateSoftwareCategory 创建软件分类
func (s *SoftwareService) CreateSoftwareCategory(params request.CreateSoftwareCategoryReq) error {
	return global.DB.Create(&models.SoftwareCategory{
		Name: params.Name,
		Sort: params.Sort,
	}).Error
}

// UpdateSoftwareCategory 更新软件分类
func (s *SoftwareService) UpdateSoftwareCategory(params request.UpdateSoftwareCategoryReq) error {
	return global.DB.Model(&models.SoftwareCategory{}).Where("id = ?", params.ID).Updates(models.SoftwareCategory{
		Name: params.Name,
		Sort: params.Sort,
	}).Error
}

// DeleteSoftwareCategory 删除软件分类
func (s *SoftwareService) DeleteSoftwareCategory(params request.DeleteSoftwareCategoryReq) error {
	if len(params.IdList) == 0 {
		return nil
	}
	return global.DB.Where("id in ?", params.IdList).Delete(&models.SoftwareCategory{}).Error
}

// ==================== Software 软件资源 CRUD ====================

// GetSoftwareList 获取软件列表
func (s *SoftwareService) GetSoftwareList(params request.GetSoftwareReq) (*global.PaginatorResp[models.Software], error) {
	query := global.DB.Model(&models.Software{})
	if params.CategoryID != 0 {
		query = query.Where("category_id = ?", params.CategoryID)
	}
	if params.Name != "" {
		query = query.Where("name LIKE ? OR version LIKE ?", "%"+params.Name+"%", "%"+params.Name+"%")
	}
	query = query.Order("created_at DESC")
	result, err := utils.Paginator[models.Software](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSoftwareById 根据ID获取软件详情
func (s *SoftwareService) GetSoftwareById(id uint) (models.Software, error) {
	var software models.Software
	err := global.DB.First(&software, id).Error
	return software, err
}

// CreateSoftware 创建软件资源
func (s *SoftwareService) CreateSoftware(params request.CreateSoftwareReq) error {
	return global.DB.Create(&models.Software{
		Name:        params.Name,
		Description: params.Description,
		Version:     params.Version,
		FileSize:    params.FileSize,
		FileURL:     params.FileURL,
		IsHot:       params.IsHot,
		CategoryID:  params.CategoryID,
	}).Error
}

// UpdateSoftware 更新软件资源
func (s *SoftwareService) UpdateSoftware(params request.UpdateSoftwareReq) error {
	return global.DB.Where("id = ?", params.ID).Updates(models.Software{
		Name:        params.Name,
		Description: params.Description,
		Version:     params.Version,
		FileSize:    params.FileSize,
		FileURL:     params.FileURL,
		IsHot:       params.IsHot,
		CategoryID:  params.CategoryID,
	}).Error
}

// DeleteSoftware 删除软件资源
func (s *SoftwareService) DeleteSoftware(params request.DeleteSoftwareReq) error {
	return global.DB.Where("id in ?", params.IdList).Delete(&models.Software{}).Error
}
