package request

import (
	"github.com/Smash249/xenova/backend/internal/global"
)

type GetProductSeriesReq struct {
	global.PaginateReq
	Name string `query:"name" json:"name"`
}

type CreateProductSeriesReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateProductSeriesReq struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type DeleteProductSeriesReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}

type GetProductReq struct {
	global.PaginateReq
	Name     string `query:"name" json:"name"`
	SeriesID uint   `query:"series_id" json:"series_id"`
	IsHot    bool   `query:"is_hot" json:"is_hot"`
}

type CreateProductReq struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Cover       string   `json:"cover" validate:"required"`
	Previews    []string `json:"previews" validate:"required"`
	SeriesID    uint     `json:"series_id" validate:"required"`
	IsHot       bool     `json:"is_hot" validate:"required"`
}

type UpdateProductReq struct {
	ID          uint     `json:"id" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Cover       string   `json:"cover" validate:"required"`
	Previews    []string `json:"previews" validate:"required"`
	SeriesID    uint     `json:"series_id" validate:"required"`
	IsHot       bool     `json:"is_hot" validate:"required"`
}

type UpdateProductHotStatusReq struct {
	ID    uint `json:"id" validate:"required"`
	IsHot bool `json:"is_hot" validate:"required"`
}

type DeleteProductReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}
