package request

import "github.com/Smash249/xenova/backend/internal/global"

type CreateSoftwareCategoryReq struct {
	Name string `json:"name" binding:"required"`
	Sort uint   `json:"sort"`
}

type UpdateSoftwareCategoryReq struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Sort uint   `json:"sort"`
}

type DeleteSoftwareCategoryReq struct {
	IdList []uint `json:"id_list" binding:"required"`
}

type GetSoftwareReq struct {
	global.PaginateReq
	CategoryID uint   `form:"category_id"`
	Name       string `form:"name"`
}

type CreateSoftwareReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Version     string `json:"version" binding:"required"`
	FileSize    string `json:"file_size" binding:"required"`
	FileURL     string `json:"file_url" binding:"required"`
	IsHot       bool   `json:"is_hot"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}

type UpdateSoftwareReq struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Version     string `json:"version" binding:"required"`
	FileSize    string `json:"file_size" binding:"required"`
	FileURL     string `json:"file_url" binding:"required"`
	IsHot       bool   `json:"is_hot"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}

type DeleteSoftwareReq struct {
	IdList []uint `json:"id_list" binding:"required"`
}
