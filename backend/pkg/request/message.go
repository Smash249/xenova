package request

import "github.com/Smash249/xenova/backend/internal/global"

type GetSystemMessageListReq struct {
	global.PaginateReq
	Name  string `query:"name"`
	Phone string `query:"phone"`
}

type CreateSystemMessageReq struct {
	Name    string `json:"name" validate:"required,min=2,max=50"`
	Content string `json:"content" validate:"required,min=5"`
	Phone   string `json:"phone" validate:"required,min=5,max=20"`
	Company string `json:"company" `
	Email   string `json:"email"`
}

type UpdateSystemMessageReq struct {
	ID      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required,min=2,max=50"`
	Content string `json:"content" validate:"required,min=5"`
	Phone   string `json:"phone" validate:"required,min=5,max=20"`
	Company string `json:"company" `
	Email   string `json:"email"`
}

type DeleteSystemMessageReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}
