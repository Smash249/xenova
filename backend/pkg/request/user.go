package request

type UserLoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type UserRegisterReq struct {
	UserName string `json:"userName" validate:"required,min=2,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Code     string `json:"code" validate:"required" `
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type UpdateUserInfoReq struct {
	UserName string `json:"userName" validate:"required,min=2,max=20"`
	Phone    string `json:"phone" validate:"required,phone"`
}

type ChangePasswordReq struct {
	OldPassword string `json:"oldPassword" validate:"required,min=8,max=20"`
	NewPassword string `json:"newPassword" validate:"required,min=8,max=20"`
}
