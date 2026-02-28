package utils

import (
	"context"
	crand "crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/Smash249/xenova/backend/internal/global"
)

const (
	codeLength     = 6
	codeExpiration = 2 * time.Minute
	codeKeyPrefix  = "email_code:"
	codeCharset    = "0123456789"
)

func GenerateVerificationCode(length int) string {
	code := make([]byte, length)
	for i := range code {
		n, _ := crand.Int(crand.Reader, big.NewInt(int64(len(codeCharset))))
		code[i] = codeCharset[n.Int64()]
	}
	return string(code)
}

func StoreAndSendEmailCode(ctx context.Context, to string) error {
	code := GenerateVerificationCode(codeLength)
	key := codeKeyPrefix + to
	if err := global.RedisClient.Set(ctx, key, code, codeExpiration).Err(); err != nil {
		return fmt.Errorf("存储验证码失败: %w", err)
	}
	content := fmt.Sprintf("您的验证码是: %s，有效期为 %d 分钟", code, int(codeExpiration.Minutes()))
	if err := global.EmailSender.Send(to, "验证码", content); err != nil {
		global.RedisClient.Del(ctx, key)
		return fmt.Errorf("发送验证码邮件失败: %w", err)
	}
	return nil
}

func VerifyEmailCode(ctx context.Context, email, code string) bool {
	key := codeKeyPrefix + email
	storedCode, err := global.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	if storedCode != code {
		return false
	}
	global.RedisClient.Del(ctx, key)
	return true
}
