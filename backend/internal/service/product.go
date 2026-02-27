package service

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/pkg/response"
	"github.com/Smash249/xenova/backend/utils"
)

type ProductService struct{}

// ==================== ProductSeries 产品系列 CRUD ====================

// GetProductSeriesList 获取产品系列列表
func (p *ProductService) GetProductSeriesList(params request.GetProductSeriesReq) (*global.PaginatorResp[models.ProductSeries], error) {
	query := global.DB.Model(&models.ProductSeries{})
	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	result, err := utils.Paginator[models.ProductSeries](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// CreateProductSeries 创建产品系列
func (p *ProductService) CreateProductSeries(params request.CreateProductSeriesReq) error {
	return global.DB.Create(&models.ProductSeries{
		Name:        params.Name,
		Description: params.Description,
	}).Error
}

// UpdateProductSeries 更新产品系列
func (p *ProductService) UpdateProductSeries(params request.UpdateProductSeriesReq) error {
	return global.DB.Model(&models.ProductSeries{}).Where("id = ?", params.ID).Updates(models.ProductSeries{
		Name:        params.Name,
		Description: params.Description,
	}).Error
}

// DeleteProductSeries 删除产品系列
func (p *ProductService) DeleteProductSeries(params request.DeleteProductSeriesReq) error {
	if len(params.IdList) == 0 {
		return nil
	}
	return global.DB.Where("id in ?", params.IdList).Delete(&models.ProductSeries{}).Error
}

// ==================== Product 产品 CRUD ====================

// GetProductList 获取产品列表
func (p *ProductService) GetProductList(params request.GetProductReq) (*global.PaginatorResp[models.Product], error) {
	query := global.DB.Model(&models.Product{})
	if params.SeriesID != 0 {
		query = query.Where("series_id = ?", params.SeriesID)
	}
	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	result, err := utils.Paginator[models.Product](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetProductDetail 根据ID获取产品
func (p *ProductService) GetProductDetail(id uint) (response.ProductDetailResp, error) {
	var resp response.ProductDetailResp
	err := global.DB.Table("product AS pd").
		Select(`
        pd.id,
        pd.series_id,
        ps.name AS series_name,
        pd.name,
        pd.description,
        pd.cover,
        pd.previews,
        pd.price,
        pd.created_at,
        pd.updated_at
    `).
		Joins("LEFT JOIN product_series AS ps ON pd.series_id = ps.id").
		Where("pd.id = ?", id).
		Scan(&resp).Error
	return resp, err
}

// CreateProduct 创建产品
func (p *ProductService) CreateProduct(params request.CreateProductReq) error {
	return global.DB.Create(&models.Product{
		Name:        params.Name,
		Description: params.Description,
		Cover:       params.Cover,
		Previews:    params.Previews,
		SeriesID:    params.SeriesID,
	}).Error
}

// UpdateProduct 更新产品
func (p *ProductService) UpdateProduct(params request.UpdateProductReq) error {
	return global.DB.Where("id = ?", params.ID).Updates(models.Product{
		Name:        params.Name,
		Description: params.Description,
		Cover:       params.Cover,
		Previews:    params.Previews,
		SeriesID:    params.SeriesID,
	}).Error
}

// DeleteProduct 删除产品
func (p *ProductService) DeleteProduct(params request.DeleteProductReq) error {
	return global.DB.Where("id in ?", params.IdList).Delete(&models.Product{}).Error
}
