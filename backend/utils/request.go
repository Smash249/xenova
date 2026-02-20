package utils

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v5"
)

func BindAndValidate[T any](ctx *echo.Context) (T, error) {
	var reqParams T
	var zero T

	// 绑定 JSON
	if err := ctx.Bind(&reqParams); err != nil {
		return zero, fmt.Errorf("参数绑定失败: %w", err)
	}

	// 校验
	if err := ctx.Validate(&reqParams); err != nil {
		return zero, fmt.Errorf("参数校验失败: %w", err)
	}

	return reqParams, nil
}

func ParseStringToUint(s string) (uint, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}
