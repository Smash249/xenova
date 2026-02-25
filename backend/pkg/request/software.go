package request

import "github.com/Smash249/xenova/backend/internal/global"

type GetSoftwareSeriesReq struct {
	global.PaginateReq
	Name string `query:"name"`
}

type CreateSoftwareSeriesReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `query:"description" validate:"required"`
}

type UpdateSoftwareSeriesReq struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type DeleteSoftwareSeriesReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}

type GetSoftwareReq struct {
	global.PaginateReq
	SeriesID uint   `query:"series_id"`
	Name     string `query:"name"`
}

type CreateSoftwareReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	FileType    string `json:"file_type" validate:"required"`
	FileSize    string `json:"file_size" validate:"required"`
	FileURL     string `json:"file_url" validate:"required"`
	IsHot       bool   `json:"is_hot" validate:"required"`
	SeriesID    uint   `json:"series_id" validate:"required"`
}

type UpdateSoftwareReq struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	FileType    string `json:"file_type" validate:"required"`
	FileSize    string `json:"file_size" validate:"required"`
	FileURL     string `json:"file_url" validate:"required"`
	IsHot       bool   `json:"is_hot" validate:"required"`
	SeriesID    uint   `json:"series_id" validate:"required"`
}

type DeleteSoftwareReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}
