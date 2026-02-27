package api

import (
	"errors"
	"net/http"

	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

var AccessoryApi = new(accessoryApi)

type accessoryApi struct{}

func (p *accessoryApi) GetAccessorySeries(ctx *echo.Context) error {
	getAccessorySeriesReq, err := utils.BindAndValidate[request.GetAccessorySeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := accessoryServiceApp.GetAccessorySeriesList(getAccessorySeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (p *accessoryApi) CreateAccessorySeries(ctx *echo.Context) error {
	createSeriesReq, err := utils.BindAndValidate[request.CreateAccessorySeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = accessoryServiceApp.CreateAccessorySeries(createSeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", http.StatusOK)
}

func (p *accessoryApi) UpdateAccessorySeries(ctx *echo.Context) error {
	updateSeriesReq, err := utils.BindAndValidate[request.UpdateAccessorySeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = accessoryServiceApp.UpdateAccessorySeries(updateSeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", http.StatusOK)
}

func (p *accessoryApi) DeleteAccessorySeries(ctx *echo.Context) error {
	deleteSeriesReq, err := utils.BindAndValidate[request.DeleteAccessorySeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = accessoryServiceApp.DeleteAccessorySeries(deleteSeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "删除成功", http.StatusOK)
}

func (p *accessoryApi) GetAccessoryList(ctx *echo.Context) error {
	getAccessoryListReq, err := utils.BindAndValidate[request.GetAccessoryReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := accessoryServiceApp.GetAccessoryList(getAccessoryListReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (p *accessoryApi) GetAccessoryDetail(ctx *echo.Context) error {
	AccessoryId := ctx.Param("accessory_id")
	if AccessoryId == "" {
		return utils.ErrorApiResponse(ctx, "配件ID不能为空", http.StatusBadRequest)
	}
	AccessoryIdUint, err := utils.ParseStringToUint(AccessoryId)
	if err != nil {
		return utils.ErrorApiResponse(ctx, "无效的产品ID", http.StatusBadRequest)
	}
	result, err := accessoryServiceApp.GetAccessoryDetail(AccessoryIdUint)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorApiResponse(ctx, "产品不存在", 404)
		}
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (p *accessoryApi) CreateAccessory(ctx *echo.Context) error {
	createAccessoryReq, err := utils.BindAndValidate[request.CreateAccessoryReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = accessoryServiceApp.CreateAccessory(createAccessoryReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", http.StatusOK)
}

func (p *accessoryApi) UpdateAccessory(ctx *echo.Context) error {
	updateAccessoryReq, err := utils.BindAndValidate[request.UpdateAccessoryReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = accessoryServiceApp.UpdateAccessory(updateAccessoryReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", http.StatusOK)
}

func (p *accessoryApi) DeleteAccessory(ctx *echo.Context) error {
	deleteAccessoryReq, err := utils.BindAndValidate[request.DeleteAccessoryReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	err = accessoryServiceApp.DeleteAccessory(deleteAccessoryReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	return utils.SuccessApiResponse(ctx, "删除成功", http.StatusOK)
}
