package initialize

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/pkg/email"
)

func InitEmail() {
	global.EmailSender = email.NewEmailSender()
}
