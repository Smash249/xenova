package request

type UserLoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type UserRegisterReq struct {
	UserName string `json:"UserName" validate:"required,min=2,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Code     string `json:"code" validate:"required" `
	Password string `json:"password" validate:"required,min=8,max=20"`
}
