package service

var UserServicesApp = new(UserService)

type UserService struct {
}

func (UserService) Login() string {
	return "hello word"
}
