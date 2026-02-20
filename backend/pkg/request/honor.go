package request

import (
	"time"

	"github.com/Smash249/xenova/backend/internal/global"
)

type GetCompanyHonorListReq struct {
	global.PaginateReq
}

type CreateCompanyHonorReq struct {
	Title       string    `json:"title"  validate:"required" `
	IssueDate   time.Time `json:"issue_date" validate:"required"`
	CertNo      string    `json:"cert_no" validate:"required"`
	Issuer      string    `json:"issuer" validate:"required"`
	Image       string    `json:"image" validate:"required"`
	Description string    `json:"description" validate:"required"`
}

type UpdateCompanyHonorReq struct {
	ID          uint      `json:"id" validate:"required"`
	Title       string    `json:"title"  validate:"required" `
	IssueDate   time.Time `json:"issue_date" validate:"required"`
	CertNo      string    `json:"cert_no" validate:"required"`
	Issuer      string    `json:"issuer" validate:"required"`
	Image       string    `json:"image" validate:"required"`
	Description string    `json:"description" validate:"required"`
}

type DeleteCompanyHonorReq struct {
	IdList []uint `json:"id_list" binding:"required"`
}

type GetLoveActivityListReq struct {
	global.PaginateReq
}

type CreateLoveActivityReq struct {
	Title        string    `json:"title"  validate:"required" `
	Location     string    `json:"location" validate:"required"`
	Participants int       `json:"participants" validate:"required"`
	Cover        string    `json:"cover" validate:"required"`
	Summary      string    `json:"summary" validate:"required"`
	Content      string    `json:"content" validate:"required"`
	ActivityDate time.Time `json:"activity_date" validate:"required"`
}

type UpdateLoveActivityReq struct {
	ID           uint      `json:"id" validate:"required"`
	Title        string    `json:"title"  validate:"required" `
	Location     string    `json:"location" validate:"required"`
	Participants int       `json:"participants" validate:"required"`
	Cover        string    `json:"cover" validate:"required"`
	Summary      string    `json:"summary" validate:"required"`
	Content      string    `json:"content" validate:"required"`
	ActivityDate time.Time `json:"activity_date" validate:"required"`
}

type DeleteLoveActivityReq struct {
	IdList []uint `json:"id_list" binding:"required"`
}

type GetCompanyPatentListReq struct {
	global.PaginateReq
}

type CreateCompanyPatentReq struct {
	Title    string    `json:"title"  validate:"required" `
	Type     string    `json:"type" validate:"required"`
	PatentNo string    `json:"patent_no" validate:"required"`
	AuthDate time.Time `json:"auth_date" validate:"required"`
	Inventor string    `json:"inventor" validate:"required"`
	Image    string    `json:"image" validate:"required"`
	Summary  string    `json:"summary" validate:"required"`
}

type UpdateCompanyPatentReq struct {
	ID       uint      `json:"id" validate:"required"`
	Title    string    `json:"title"  validate:"required" `
	Type     string    `json:"type" validate:"required"`
	PatentNo string    `json:"patent_no" validate:"required"`
	AuthDate time.Time `json:"auth_date" validate:"required"`
	Inventor string    `json:"inventor" validate:"required"`
	Image    string    `json:"image" validate:"required"`
	Summary  string    `json:"summary" validate:"required"`
}

type DeleteCompanyPatentReq struct {
	IdList []uint `json:"id_list" binding:"required"`
}
