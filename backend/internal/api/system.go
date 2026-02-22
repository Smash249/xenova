package api

import (
	"net/http"

	"github.com/Smash249/xenova/backend/utils"
	"github.com/labstack/echo/v5"
)

var SystemApi = new(systemApi)

type systemApi struct{}

func (systemApi) Health(ctx *echo.Context) error {
	return utils.SuccessApiResponse(ctx, "ok", http.StatusOK)
}
