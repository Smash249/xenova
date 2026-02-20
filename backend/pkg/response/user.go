package response

import "github.com/Smash249/xenova/backend/internal/models"

type UserLoginResp struct {
	User         models.User `json:"user"`
	AccessToken  string      `json:"accessToken"`
	ReFreshToken string      `json:"reFreshToken"`
}
