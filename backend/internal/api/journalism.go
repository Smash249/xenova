package api

import (
	"errors"
	"net/http"

	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

var JournalismApi = new(journalismApi)

type journalismApi struct{}

func (j *journalismApi) GetJournalismSeries(ctx *echo.Context) error {
	result, err := journalismServiceApp.GetJournalismSeriesList()
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (j *journalismApi) CreateJournalismSeries(ctx *echo.Context) error {
	createSeriesReq, err := utils.BindAndValidate[request.CreateJournalismSeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = journalismServiceApp.CreateJournalismSeries(createSeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", http.StatusOK)
}

func (j *journalismApi) UpdateJournalismSeries(ctx *echo.Context) error {
	updateSeriesReq, err := utils.BindAndValidate[request.UpdateJournalismSeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = journalismServiceApp.UpdateJournalismSeries(updateSeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", http.StatusOK)
}

func (j *journalismApi) DeleteJournalismSeries(ctx *echo.Context) error {
	deleteSeriesReq, err := utils.BindAndValidate[request.DeleteJournalismSeriesReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = journalismServiceApp.DeleteJournalismSeries(deleteSeriesReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "删除成功", http.StatusOK)
}

func (j *journalismApi) GetJournalismList(ctx *echo.Context) error {
	getJournalismListReq, err := utils.BindAndValidate[request.GetJournalismReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	result, err := journalismServiceApp.GetJournalismList(getJournalismListReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (j *journalismApi) GetJournalismDetail(ctx *echo.Context) error {
	journalismId := ctx.Param("journalism_id")
	if journalismId == "" {
		return utils.ErrorApiResponse(ctx, "新闻ID不能为空", http.StatusBadRequest)
	}
	journalismIdUint, err := utils.ParseStringToUint(journalismId)
	if err != nil {
		return utils.ErrorApiResponse(ctx, "无效的新闻ID", http.StatusBadRequest)
	}
	result, err := journalismServiceApp.GetJournalismById(journalismIdUint)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorApiResponse(ctx, "新闻不存在", 404)
		}
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}

func (j *journalismApi) CreateJournalism(ctx *echo.Context) error {
	createJournalismReq, err := utils.BindAndValidate[request.CreateJournalismReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = journalismServiceApp.CreateJournalism(createJournalismReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "创建成功", http.StatusOK)
}

func (j *journalismApi) UpdateJournalism(ctx *echo.Context) error {
	updateJournalismReq, err := utils.BindAndValidate[request.UpdateJournalismReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = journalismServiceApp.UpdateJournalism(updateJournalismReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	return utils.SuccessApiResponse(ctx, "更新成功", http.StatusOK)
}

func (j *journalismApi) DeleteJournalism(ctx *echo.Context) error {
	deleteJournalismReq, err := utils.BindAndValidate[request.DeleteJournalismReq](ctx)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	err = journalismServiceApp.DeleteJournalism(deleteJournalismReq)
	if err != nil {
		return utils.ErrorApiResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	return utils.SuccessApiResponse(ctx, "删除成功", http.StatusOK)
}
