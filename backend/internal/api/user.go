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
		return utils.ErroApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := userservicesApp.Login(userLoginReq)
	if err != nil {
		return utils.ErroApiResponse(ctx, err, http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (userApi) Register(ctx *echo.Context) error {
	userRegisterReq, err := utils.BindAndValidate[request.UserRegisterReq](ctx)
	if err != nil {
		return utils.ErroApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	if err := userservicesApp.Register(userRegisterReq); err != nil {
		return utils.ErroApiResponse(ctx, err, http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "注册成功", http.StatusOK)
}
