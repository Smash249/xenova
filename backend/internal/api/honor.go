package api

import (
	"net/http"

	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
	"github.com/labstack/echo/v5"
)

var HonorApi = new(honorApi)

type honorApi struct{}

func (h *honorApi) GetCompanyHonorList(ctx *echo.Context) error {
	getCompanyHonorListReq, err := utils.BindAndValidate[request.GetCompanyHonorListReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := honorServiceApp.GetCompanyHonorList(getCompanyHonorListReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (h *honorApi) CreateCompanyHonor(ctx *echo.Context) error {
	createCompanyHonorReq, err := utils.BindAndValidate[request.CreateCompanyHonorReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = honorServiceApp.CreateCompanyHonor(createCompanyHonorReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", http.StatusOK)
}

func (h *honorApi) UpdateCompanyHonor(ctx *echo.Context) error {
	updateCompanyHonorReq, err := utils.BindAndValidate[request.UpdateCompanyHonorReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = honorServiceApp.UpdateCompanyHonor(updateCompanyHonorReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", http.StatusOK)
}

func (h *honorApi) DeleteCompanyHonor(ctx *echo.Context) error {
	deleteCompanyHonorReq, err := utils.BindAndValidate[request.DeleteCompanyHonorReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = honorServiceApp.DeleteCompanyHonor(deleteCompanyHonorReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "删除成功", http.StatusOK)
}

func (h *honorApi) GetLoveActivityList(ctx *echo.Context) error {
	getLoveActivityListReq, err := utils.BindAndValidate[request.GetLoveActivityListReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := honorServiceApp.GetActivityList(getLoveActivityListReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (h *honorApi) CreateLoveActivity(ctx *echo.Context) error {
	createLoveActivityReq, err := utils.BindAndValidate[request.CreateLoveActivityReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = honorServiceApp.CreateActivity(createLoveActivityReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", http.StatusOK)
}

func (h *honorApi) UpdateLoveActivity(ctx *echo.Context) error {
	updateLoveActivityReq, err := utils.BindAndValidate[request.UpdateLoveActivityReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = honorServiceApp.UpdateActivity(updateLoveActivityReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", http.StatusOK)
}

func (h *honorApi) DeleteLoveActivity(ctx *echo.Context) error {
	deleteLoveActivityReq, err := utils.BindAndValidate[request.DeleteLoveActivityReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = honorServiceApp.DeleteActivity(deleteLoveActivityReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "删除成功", http.StatusOK)
}

func (h *honorApi) GetCompanyPatentList(ctx *echo.Context) error {
	getCompanyPatentListReq, err := utils.BindAndValidate[request.GetCompanyPatentListReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := honorServiceApp.GetPatentList(getCompanyPatentListReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (h *honorApi) CreateCompanyPatent(ctx *echo.Context) error {
	createCompanyPatentReq, err := utils.BindAndValidate[request.CreateCompanyPatentReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = honorServiceApp.CreatePatent(createCompanyPatentReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", http.StatusOK)
}

func (h *honorApi) UpdateCompanyPatent(ctx *echo.Context) error {
	updateCompanyPatentReq, err := utils.BindAndValidate[request.UpdateCompanyPatentReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = honorServiceApp.UpdatePatent(updateCompanyPatentReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", http.StatusOK)
}

func (h *honorApi) DeleteCompanyPatent(ctx *echo.Context) error {
	deleteCompanyPatentReq, err := utils.BindAndValidate[request.DeleteCompanyPatentReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = honorServiceApp.DeletePatent(deleteCompanyPatentReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "删除成功", http.StatusOK)
}
