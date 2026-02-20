package utils

import "github.com/labstack/echo/v5"

func SuccessApiResponse(ctx *echo.Context, data any, code int) error {
	return ctx.JSON(code, map[string]any{
		"success": true,
		"data":    data,
	})
}

func ErrorApiResponse(ctx *echo.Context, err any, code int) error {
	return ctx.JSON(code, map[string]any{
		"message": err,
		"success": false,
	})
}
