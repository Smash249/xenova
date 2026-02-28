package api

import (
	"net/http"

	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
	"github.com/labstack/echo/v5"
)

var UserApi = new(userApi)

type userApi struct{}

func (userApi) FrontendLogin(ctx *echo.Context) error {
	userLoginReq, err := utils.BindAndValidate[request.UserLoginReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := userServicesApp.FrontendLogin(userLoginReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (userApi) AdminLogin(ctx *echo.Context) error {
	userLoginReq, err := utils.BindAndValidate[request.UserLoginReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := userServicesApp.AdminLogin(userLoginReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (userApi) SendEmailCode(ctx *echo.Context) error {
	emailCodeReq, err := utils.BindAndValidate[request.EmailCodeReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	if err := userServicesApp.SendEmailCode(emailCodeReq); err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "验证码发送成功", http.StatusOK)
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

func (userApi) GetUserInfo(ctx *echo.Context) error {
	userId, exists := utils.GetUserID(ctx)
	if !exists {
		return utils.ErrorApiResponse(ctx, "获取用户失败", http.StatusBadRequest)
	}
	result, err := userServicesApp.GetUserInfo(userId)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (userApi) UpdateUserInfo(ctx *echo.Context) error {
	userId, exists := utils.GetUserID(ctx)
	if !exists {
		return utils.ErrorApiResponse(ctx, "获取用户失败", http.StatusBadRequest)
	}
	userUpdateInfoReq, err := utils.BindAndValidate[request.UpdateUserInfoReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := userServicesApp.UpdateUserInfo(userId, userUpdateInfoReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (userApi) AdminUpdateUserInfo(ctx *echo.Context) error {
	userId, exists := utils.GetUserID(ctx)
	if !exists {
		return utils.ErrorApiResponse(ctx, "获取用户失败", http.StatusBadRequest)
	}
	updateAdminInfoReq, err := utils.BindAndValidate[request.UpdateAdminInfoReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := userServicesApp.AdminUpdateUserInfo(userId, updateAdminInfoReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (userApi) ChangePassword(ctx *echo.Context) error {
	userId, exists := utils.GetUserID(ctx)
	if !exists {
		return utils.ErrorApiResponse(ctx, "获取用户失败", http.StatusBadRequest)
	}
	changePasswordReq, err := utils.BindAndValidate[request.ChangePasswordReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	if err := userServicesApp.ChangePassword(userId, changePasswordReq); err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "密码修改成功", http.StatusOK)
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
