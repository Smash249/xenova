package request

import (
	"testing"
)

func TestUserLoginReq(t *testing.T) {
	tests := []struct {
		name  string
		email string
		pwd   string
	}{
		{
			name:  "Valid login request",
			email: "test@example.com",
			pwd:   "password123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := UserLoginReq{
				Email:    tt.email,
				Password: tt.pwd,
			}
			if req.Email != tt.email {
				t.Errorf("Expected email %s, got %s", tt.email, req.Email)
			}
			if req.Password != tt.pwd {
				t.Errorf("Expected password %s, got %s", tt.pwd, req.Password)
			}
		})
	}
}

func TestUserRegisterReq(t *testing.T) {
	tests := []struct {
		name     string
		username string
		email    string
		code     string
		password string
	}{
		{
			name:     "Valid register request",
			username: "johndoe",
			email:    "john@example.com",
			code:     "123456",
			password: "password123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := UserRegisterReq{
				UserName: tt.username,
				Email:    tt.email,
				Code:     tt.code,
				Password: tt.password,
			}
			if req.UserName != tt.username {
				t.Errorf("Expected username %s, got %s", tt.username, req.UserName)
			}
			if req.Email != tt.email {
				t.Errorf("Expected email %s, got %s", tt.email, req.Email)
			}
		})
	}
}
