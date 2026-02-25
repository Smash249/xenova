package middleware

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

// NewAdmin 仅允许 admin 角色访问
func NewAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx *echo.Context) error {
			role := ctx.Get("role")
			if role == nil {
				return echo.NewHTTPError(http.StatusForbidden, "无权限操作")
			}
			if roleStr, ok := role.(string); !ok || roleStr != "admin" {
				return echo.NewHTTPError(http.StatusForbidden, "无权限操作")
			}
			return next(ctx)
		}
	}
}
