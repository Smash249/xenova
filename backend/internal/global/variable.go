package global

import (
	"github.com/Smash249/xenova/backend/pkg/email"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client
var EmailSender *email.EmailSender
