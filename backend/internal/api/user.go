package api

import "github.com/labstack/echo/v5"

var UserApi = new(userApi)

type userApi struct{}

func (a *userApi) GetUserInfo(ctx *echo.Context) error {
	return ctx.String(200, "User Info")
}
