package request

import (
	"github.com/Smash249/xenova/backend/internal/global"
)

type GetAccessorySeriesReq struct {
	global.PaginateReq
	Name string `query:"name" json:"name"`
}

type CreateAccessorySeriesReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateAccessorySeriesReq struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type DeleteAccessorySeriesReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}

type GetAccessoryReq struct {
	global.PaginateReq
	Name     string `query:"name" json:"name"`
	SeriesID uint   `query:"series_id" json:"series_id"`
}

type CreateAccessoryReq struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Cover       string   `json:"cover" validate:"required"`
	Previews    []string `json:"previews" validate:"required"`
	SeriesID    uint     `json:"series_id" validate:"required"`
	Price       float64  `json:"price" validate:"required"`
}

type UpdateAccessoryReq struct {
	ID          uint     `json:"id" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Cover       string   `json:"cover" validate:"required"`
	Previews    []string `json:"previews" validate:"required"`
	SeriesID    uint     `json:"series_id" validate:"required"`
	Price       float64  `json:"price" validate:"required"`
}

type DeleteAccessoryReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}
