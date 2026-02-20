package service

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
)

type SoftwareService struct{}

// ==================== SoftwareSeries 软件分类 CRUD ====================

// GetSoftwareSeriesList 获取软件分类列表
func (s *SoftwareService) GetSoftwareSeriesList() ([]models.SoftwareSeries, error) {
	var categoryList []models.SoftwareSeries
	err := global.DB.Find(&categoryList).Error
	return categoryList, err
}

// CreateSoftwareSeries 创建软件分类
func (s *SoftwareService) CreateSoftwareSeries(params request.CreateSoftwareSeriesReq) error {
	return global.DB.Create(&models.SoftwareSeries{
		Name: params.Name,
	}).Error
}

// UpdateSoftwareSeries 更新软件分类
func (s *SoftwareService) UpdateSoftwareSeries(params request.UpdateSoftwareSeriesReq) error {
	return global.DB.Model(&models.SoftwareSeries{}).Where("id = ?", params.ID).Updates(models.SoftwareSeries{
		Name: params.Name,
	}).Error
}

// DeleteSoftwareSeries 删除软件分类
func (s *SoftwareService) DeleteSoftwareSeries(params request.DeleteSoftwareSeriesReq) error {
	if len(params.IdList) == 0 {
		return nil
	}
	return global.DB.Where("id in ?", params.IdList).Delete(&models.SoftwareSeries{}).Error
}

// ==================== Software 软件资源 CRUD ====================

// GetSoftwareList 获取软件列表
func (s *SoftwareService) GetSoftwareList(params request.GetSoftwareReq) (*global.PaginatorResp[models.Software], error) {
	query := global.DB.Model(&models.Software{})
	if params.SeriesID != 0 {
		query = query.Where("series_id = ?", params.SeriesID)
	}
	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
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
		FileSize:    params.FileSize,
		FileURL:     params.FileURL,
		FileType:    params.FileType,
		IsHot:       params.IsHot,
		SeriesID:    params.SeriesID,
	}).Error
}

// UpdateSoftware 更新软件资源
func (s *SoftwareService) UpdateSoftware(params request.UpdateSoftwareReq) error {
	return global.DB.Where("id = ?", params.ID).Updates(models.Software{
		Name:        params.Name,
		Description: params.Description,
		FileSize:    params.FileSize,
		FileURL:     params.FileURL,
		FileType:    params.FileType,
		IsHot:       params.IsHot,
		SeriesID:    params.SeriesID,
	}).Error
}

// DeleteSoftware 删除软件资源
func (s *SoftwareService) DeleteSoftware(params request.DeleteSoftwareReq) error {
	return global.DB.Where("id in ?", params.IdList).Delete(&models.Software{}).Error
}
