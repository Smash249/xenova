package request

import "github.com/Smash249/xenova/backend/internal/global"

type GetSystemJobPositionListReq struct {
	global.PaginateReq
	Title      string `query:"title" form:"title"`
	Department string `query:"department" form:"department"`
	Status     int    `query:"status" form:"status"`
}

type CreateSystemJobPositionReq struct {
	Title            string `json:"title" validate:"required,min=2,max=100"`
	Department       string `json:"department" validate:"required"`
	Location         string `json:"location" validate:"required"`
	Headcount        string `json:"headcount" validate:"required"`
	Experience       string `json:"experience" validate:"required"`
	Education        string `json:"education" validate:"required"`
	SalaryRange      string `json:"salaryRange" validate:"required"`
	Responsibilities string `json:"responsibilities" validate:"required"`
	Requirements     string `json:"requirements" validate:"required"`
}

type UpdateSystemJobPositionReq struct {
	ID               uint   `json:"id" validate:"required"`
	Title            string `json:"title" validate:"required,min=2,max=100"`
	Department       string `json:"department" validate:"required"`
	Location         string `json:"location" validate:"required"`
	Headcount        string `json:"headcount" validate:"required"`
	Experience       string `json:"experience" validate:"required"`
	Education        string `json:"education" validate:"required"`
	SalaryRange      string `json:"salaryRange" validate:"required"`
	Responsibilities string `json:"responsibilities" validate:"required"`
	Requirements     string `json:"requirements" validate:"required"`
	Status           int    `json:"status" validate:"required"`
	Sort             int    `json:"sort" validate:"required"`
}

type DeleteSystemJobPositionReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}
