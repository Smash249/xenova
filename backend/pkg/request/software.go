package request

import "github.com/Smash249/xenova/backend/internal/global"

type GetSoftwareSeriesReq struct {
	global.PaginateReq
	Name string `query:"name"`
}

type CreateSoftwareSeriesReq struct {
	Name string `json:"name" binding:"required"`
	Sort uint   `json:"sort"`
}

type UpdateSoftwareSeriesReq struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Sort uint   `json:"sort"`
}

type DeleteSoftwareSeriesReq struct {
	IdList []uint `json:"id_list" binding:"required"`
}

type GetSoftwareReq struct {
	global.PaginateReq
	SeriesID uint   `query:"series_id"`
	Name     string `query:"name"`
}

type CreateSoftwareReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	FileType    string `json:"file_type" binding:"required"`
	FileSize    string `json:"file_size" binding:"required"`
	FileURL     string `json:"file_url" binding:"required"`
	IsHot       bool   `json:"is_hot" binding:"required"`
	SeriesID    uint   `json:"series_id" binding:"required"`
}

type UpdateSoftwareReq struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	FileType    string `json:"file_type" binding:"required"`
	FileSize    string `json:"file_size" binding:"required"`
	FileURL     string `json:"file_url" binding:"required"`
	IsHot       bool   `json:"is_hot" binding:"required"`
	SeriesID    uint   `json:"series_id" binding:"required"`
}

type DeleteSoftwareReq struct {
	IdList []uint `json:"id_list" binding:"required"`
}
