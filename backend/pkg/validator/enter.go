package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type customValidator struct {
	validator *validator.Validate
}

func InitCustomValidator() *customValidator {
	return &customValidator{
		validator: validator.New(),
	}
}

func (cv *customValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}
