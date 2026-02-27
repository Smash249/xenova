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

func (systemApi) GetSystemMessageList(ctx *echo.Context) error {
	getSystemMessageListReq, err := utils.BindAndValidate[request.GetSystemMessageListReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := systemServiceApp.GetSystemMessageList(getSystemMessageListReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (systemApi) CreateSystemMessage(ctx *echo.Context) error {
	createSystemMessageReq, err := utils.BindAndValidate[request.CreateSystemMessageReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	userId, ok := utils.GetUserID(ctx)
	if !ok {
		return utils.ErrorApiResponse(ctx, "用户未登录", http.StatusUnauthorized)
	}
	err = systemServiceApp.CreateSystemMessage(userId, createSystemMessageReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return utils.SuccessApiResponse(ctx, "消息已创建", http.StatusOK)
}

func (systemApi) UpdateSystemMessage(ctx *echo.Context) error {
	updateSystemMessageReq, err := utils.BindAndValidate[request.UpdateSystemMessageReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = systemServiceApp.UpdateSystemMessage(updateSystemMessageReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return utils.SuccessApiResponse(ctx, "消息已更新", http.StatusOK)
}

func (systemApi) DeleteSystemMessage(ctx *echo.Context) error {
	deleteSystemMessageReq, err := utils.BindAndValidate[request.DeleteSystemMessageReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = systemServiceApp.DeleteSystemMessage(deleteSystemMessageReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return utils.SuccessApiResponse(ctx, "消息已删除", http.StatusOK)
}

func (systemApi) GetSystemJobPositionList(ctx *echo.Context) error {
	getSystemJobPostListReq, err := utils.BindAndValidate[request.GetSystemJobPositionListReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := systemServiceApp.GetSystemJobPositionList(getSystemJobPostListReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (systemApi) CreateSystemJobPosition(ctx *echo.Context) error {
	createSystemJobPositionReq, err := utils.BindAndValidate[request.CreateSystemJobPositionReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = systemServiceApp.CreateSystemJobPosition(createSystemJobPositionReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return utils.SuccessApiResponse(ctx, "职位已创建", http.StatusOK)
}

func (systemApi) UpdateSystemJobPosition(ctx *echo.Context) error {
	updateSystemJobPositionReq, err := utils.BindAndValidate[request.UpdateSystemJobPositionReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = systemServiceApp.UpdateSystemJobPosition(updateSystemJobPositionReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return utils.SuccessApiResponse(ctx, "职位已更新", http.StatusOK)
}

func (systemApi) DeleteSystemJobPosition(ctx *echo.Context) error {
	deleteSystemJobPositionReq, err := utils.BindAndValidate[request.DeleteSystemJobPositionReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = systemServiceApp.DeleteSystemJobPosition(deleteSystemJobPositionReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return utils.SuccessApiResponse(ctx, "职位已删除", http.StatusOK)
}
