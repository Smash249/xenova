package request

import (
	"github.com/Smash249/xenova/backend/internal/global"
)

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
	SeriesID uint   `query:"seriesId" json:"seriesId" validate:"required"`
}

type CreateProductReq struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Cover       string   `json:"cover" validate:"required"`
	Previews    []string `json:"previews" validate:"required"`
	Price       float64  `json:"price" validate:"required"`
	SeriesID    uint     `json:"seriesId" validate:"required"`
}

type UpdateProductReq struct {
	ID          uint     `json:"id" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Cover       string   `json:"cover" validate:"required"`
	Previews    []string `json:"previews" validate:"required"`
	Price       float64  `json:"price" validate:"required"`
	SeriesID    uint     `json:"seriesId" validate:"required"`
}

type DeleteProductReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}
