package api

import (
	"errors"

	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

var ProductApi = new(productApi)

type productApi struct{}

func (p *productApi) GetProductSeries(ctx *echo.Context) error {
	result, err := productServiceApp.GetProductSeriesList()
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	return utils.SuccessApiResponse(ctx, result, 200)
}

func (p *productApi) CreateProductSeries(ctx *echo.Context) error {
	createSeriesReq, err := utils.BindAndValidate[request.CreateProductSeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	err = productServiceApp.CreateProductSeries(createSeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", 200)
}

func (p *productApi) UpdateProductSeries(ctx *echo.Context) error {
	updateSeriesReq, err := utils.BindAndValidate[request.UpdateProductSeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	err = productServiceApp.UpdateProductSeries(updateSeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", 200)
}

func (p *productApi) DeleteProductSeries(ctx *echo.Context) error {
	deleteSeriesReq, err := utils.BindAndValidate[request.DeleteProductSeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	err = productServiceApp.DeleteProductSeries(deleteSeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	return utils.SuccessApiResponse(ctx, "删除成功", 200)
}

func (p *productApi) GetProductList(ctx *echo.Context) error {
	getProductListReq, err := utils.BindAndValidate[request.GetProductReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	result, err := productServiceApp.GetProductList(getProductListReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	return utils.SuccessApiResponse(ctx, result, 200)
}

func (p *productApi) GetProductDetail(ctx *echo.Context) error {
	productId := ctx.Param("product_id")
	if productId == "" {
		return utils.ErrorApiResponse(ctx, "产品ID不能为空", 400)
	}
	productIdUint, err := utils.ParseStringToUint(productId)
	if err != nil {
		return utils.ErrorApiResponse(ctx, "无效的产品ID", 400)
	}
	result, err := productServiceApp.GetProductById(productIdUint)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorApiResponse(ctx, "产品不存在", 404)
		}
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	return utils.SuccessApiResponse(ctx, result, 200)
}

func (p *productApi) CreateProduct(ctx *echo.Context) error {
	createProductReq, err := utils.BindAndValidate[request.CreateProductReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	err = productServiceApp.CreateProduct(createProductReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", 200)
}

func (p *productApi) UpdateProduct(ctx *echo.Context) error {
	updateProductReq, err := utils.BindAndValidate[request.UpdateProductReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	err = productServiceApp.UpdateProduct(updateProductReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", 200)
}

func (p *productApi) DeleteProduct(ctx *echo.Context) error {
	deleteProductReq, err := utils.BindAndValidate[request.DeleteProductReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}

	err = productServiceApp.DeleteProduct(deleteProductReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), 400)
	}

	return utils.SuccessApiResponse(ctx, "删除成功", 200)
}
