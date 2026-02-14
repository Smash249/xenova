package response

import (
	"testing"

	"github.com/Smash249/xenova/backend/internal/models"
)

func TestUserLoginResp(t *testing.T) {
	user := models.User{
		UserName: "testuser",
		Email:    "test@example.com",
	}

	resp := UserLoginResp{
		User:         user,
		AccessToken:  "test_access_token",
		ReFreshToken: "test_refresh_token",
	}

	if resp.User.UserName != "testuser" {
		t.Errorf("Expected username 'testuser', got %s", resp.User.UserName)
	}
	if resp.AccessToken != "test_access_token" {
		t.Errorf("Expected access token 'test_access_token', got %s", resp.AccessToken)
	}
	if resp.ReFreshToken != "test_refresh_token" {
		t.Errorf("Expected refresh token 'test_refresh_token', got %s", resp.ReFreshToken)
	}
}
