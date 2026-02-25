package api

import (
	"net/http"
	"strconv"

	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
	"github.com/labstack/echo/v5"
)

var SystemApi = new(systemApi)

type systemApi struct{}

func (systemApi) Health(ctx *echo.Context) error {
	return utils.SuccessApiResponse(ctx, "ok", http.StatusOK)
}

func (systemApi) GetSystemUserList(ctx *echo.Context) error {
	getSystemUserListReq, err := utils.BindAndValidate[request.GetSystemUserListReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := systemServiceApp.GetSystemUserList(getSystemUserListReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (systemApi) BanedUserById(ctx *echo.Context) error {
	userId := ctx.Param("id")
	if userId == "" {
		return utils.ErrorApiResponse(ctx, "用户ID不能为空", http.StatusBadRequest)
	}
	parsedId, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return utils.ErrorApiResponse(ctx, "无效的用户ID", http.StatusBadRequest)
	}
	err = systemServiceApp.BanedUserById(uint(parsedId))
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return utils.SuccessApiResponse(ctx, "用户已封禁", http.StatusOK)
}
