package api

import (
	"net/http"

	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
	"github.com/labstack/echo/v5"
)

var UserApi = new(userApi)

type userApi struct{}

func (userApi) Login(ctx *echo.Context) error {
	userLoginReq, err := utils.BindAndValidate[request.UserLoginReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := userServicesApp.Login(userLoginReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (userApi) Register(ctx *echo.Context) error {
	userRegisterReq, err := utils.BindAndValidate[request.UserRegisterReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	if err := userServicesApp.Register(userRegisterReq); err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "注册成功", http.StatusOK)
}

func (userApi) Health(ctx *echo.Context) error {
	return utils.SuccessApiResponse(ctx, "ok", http.StatusOK)
}

func (userApi) Upload(ctx *echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := userServicesApp.Upload(file)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx,
		map[string]string{
			"url": result,
		}, http.StatusOK)
}
