package service

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
)

var ProductServiceApp = new(ProductService)

type ProductService struct{}

// ==================== ProductSeries 产品系列 CRUD ====================

// GetProductSeriesList 获取产品系列列表
func (p *ProductService) GetProductSeriesList() ([]models.ProductSeries, error) {
	var seriesList []models.ProductSeries
	err := global.DB.Find(&seriesList).Error
	return seriesList, err
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
	query := global.DB.Model(&models.Product{}).Where("series_id = ?", params.SeriesID)
	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	result, err := utils.Paginator[models.Product](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetProductById 根据ID获取产品
func (p *ProductService) GetProductById(id uint) (models.Product, error) {
	var product models.Product
	err := global.DB.First(&product, id).Error
	return product, err
}

// CreateProduct 创建产品
func (p *ProductService) CreateProduct(params request.CreateProductReq) error {
	return global.DB.Create(&models.Product{
		Name:        params.Name,
		Description: params.Description,
		Cover:       params.Cover,
		Previews:    params.Previews,
		Price:       params.Price,
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
		Price:       params.Price,
		SeriesID:    params.SeriesID,
	}).Error
}

// DeleteProduct 删除产品
func (p *ProductService) DeleteProduct(params request.DeleteProductReq) error {
	return global.DB.Where("id in ?", params.IdList).Delete(&models.Product{}).Error
}
