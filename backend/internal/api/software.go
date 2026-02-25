package api

import (
	"errors"
	"net/http"

	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

var SoftwareApi = new(softwareApi)

type softwareApi struct{}

func (s *softwareApi) GetSoftwareSeries(ctx *echo.Context) error {
	getSoftwareSeriesReq, err := utils.BindAndValidate[request.GetSoftwareSeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := softwareServiceApp.GetSoftwareSeriesList(getSoftwareSeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (s *softwareApi) CreateSoftwareSeries(ctx *echo.Context) error {
	createCategoryReq, err := utils.BindAndValidate[request.CreateSoftwareSeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = softwareServiceApp.CreateSoftwareSeries(createCategoryReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", http.StatusOK)
}

func (s *softwareApi) UpdateSoftwareSeries(ctx *echo.Context) error {
	updateCategoryReq, err := utils.BindAndValidate[request.UpdateSoftwareSeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = softwareServiceApp.UpdateSoftwareSeries(updateCategoryReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", http.StatusOK)
}

func (s *softwareApi) DeleteSoftwareSeries(ctx *echo.Context) error {
	deleteCategoryReq, err := utils.BindAndValidate[request.DeleteSoftwareSeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = softwareServiceApp.DeleteSoftwareSeries(deleteCategoryReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "删除成功", http.StatusOK)
}

func (s *softwareApi) GetSoftwareList(ctx *echo.Context) error {
	getSoftwareListReq, err := utils.BindAndValidate[request.GetSoftwareReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := softwareServiceApp.GetSoftwareList(getSoftwareListReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (s *softwareApi) GetSoftwareDetail(ctx *echo.Context) error {
	softwareId := ctx.Param("software_id")
	if softwareId == "" {
		return utils.ErrorApiResponse(ctx, "软件ID不能为空", http.StatusBadRequest)
	}
	softwareIdUint, err := utils.ParseStringToUint(softwareId)
	if err != nil {
		return utils.ErrorApiResponse(ctx, "无效的软件ID", http.StatusBadRequest)
	}
	result, err := softwareServiceApp.GetSoftwareById(softwareIdUint)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorApiResponse(ctx, "软件不存在", 404)
		}
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (s *softwareApi) CreateSoftware(ctx *echo.Context) error {
	createSoftwareReq, err := utils.BindAndValidate[request.CreateSoftwareReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = softwareServiceApp.CreateSoftware(createSoftwareReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", http.StatusOK)
}

func (s *softwareApi) UpdateSoftware(ctx *echo.Context) error {
	updateSoftwareReq, err := utils.BindAndValidate[request.UpdateSoftwareReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = softwareServiceApp.UpdateSoftware(updateSoftwareReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", http.StatusOK)
}

func (s *softwareApi) DeleteSoftware(ctx *echo.Context) error {
	deleteSoftwareReq, err := utils.BindAndValidate[request.DeleteSoftwareReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = softwareServiceApp.DeleteSoftware(deleteSoftwareReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "删除成功", http.StatusOK)
}
