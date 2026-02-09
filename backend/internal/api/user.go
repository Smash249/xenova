package api

import (
	"net/http"

	"github.com/Smash249/xenova/backend/pkg/utils"
	"github.com/labstack/echo/v5"
)

var UserApi = new(userApi)

type userApi struct{}

func (userApi) Login(ctx *echo.Context) error {
	result := userservicesApp.Login()
	return utils.SuccessApiResponse(ctx, result, http.StatusOK)
}
