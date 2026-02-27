package service

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/pkg/response"
	"github.com/Smash249/xenova/backend/utils"
)

type AccessoryService struct{}

// ==================== AccessorySeries 配件系列 CRUD ====================

// GetAccessorySeriesList 获取配件系列列表
func (p *AccessoryService) GetAccessorySeriesList(params request.GetAccessorySeriesReq) (*global.PaginatorResp[models.AccessorySeries], error) {
	query := global.DB.Model(&models.AccessorySeries{})
	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	result, err := utils.Paginator[models.AccessorySeries](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// CreateAccessorySeries 配件产品系列
func (p *AccessoryService) CreateAccessorySeries(params request.CreateAccessorySeriesReq) error {
	return global.DB.Create(&models.AccessorySeries{
		Name:        params.Name,
		Description: params.Description,
	}).Error
}

// UpdateAccessorySeries 更新配件系列
func (p *AccessoryService) UpdateAccessorySeries(params request.UpdateAccessorySeriesReq) error {
	return global.DB.Model(&models.AccessorySeries{}).Where("id = ?", params.ID).Updates(models.AccessorySeries{
		Name:        params.Name,
		Description: params.Description,
	}).Error
}

// DeleteAccessorySeries 删除配件系列
func (p *AccessoryService) DeleteAccessorySeries(params request.DeleteAccessorySeriesReq) error {
	if len(params.IdList) == 0 {
		return nil
	}
	return global.DB.Where("id in ?", params.IdList).Delete(&models.AccessorySeries{}).Error
}

// ==================== Accessory 配件 CRUD ====================

// GetAccessoryList 获取配件列表
func (p *AccessoryService) GetAccessoryList(params request.GetAccessoryReq) (*global.PaginatorResp[models.Accessory], error) {
	query := global.DB.Model(&models.Accessory{})
	if params.SeriesID != 0 {
		query = query.Where("series_id = ?", params.SeriesID)
	}
	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	result, err := utils.Paginator[models.Accessory](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccessoryDetail 根据ID获取配件
func (p *AccessoryService) GetAccessoryDetail(id uint) (response.AccessoryDetailResp, error) {
	var resp response.AccessoryDetailResp
	err := global.DB.Table("accessory AS acsy").
		Select(`
        acsy.id,
        acsy.series_id,
        acsy.name AS series_name,
        acsy.name,
        acsy.description,
        acsy.cover,
        acsy.previews,
        acsy.price,
        acsy.created_at,
        acsy.updated_at
    `).
		Joins("LEFT JOIN accessory_series AS asy ON acsy.series_id = asy.id").
		Where("acsy.id = ?", id).
		Scan(&resp).Error
	return resp, err
}

// CreateAccessory 创建配件
func (p *AccessoryService) CreateAccessory(params request.CreateAccessoryReq) error {
	return global.DB.Create(&models.Accessory{
		Name:        params.Name,
		Description: params.Description,
		Cover:       params.Cover,
		Previews:    params.Previews,
		SeriesID:    params.SeriesID,
	}).Error
}

// UpdateAccessory 更新配件
func (p *AccessoryService) UpdateAccessory(params request.UpdateAccessoryReq) error {
	return global.DB.Where("id = ?", params.ID).Updates(models.Accessory{
		Name:        params.Name,
		Description: params.Description,
		Cover:       params.Cover,
		Previews:    params.Previews,
		SeriesID:    params.SeriesID,
		Price:       params.Price,
	}).Error
}

// DeleteAccessory 删除配件
func (p *AccessoryService) DeleteAccessory(params request.DeleteAccessoryReq) error {
	return global.DB.Where("id in ?", params.IdList).Delete(&models.Accessory{}).Error
}
