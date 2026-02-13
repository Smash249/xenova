package utils

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/spf13/viper"
)

// CustomClaims 自定义 JWT 声明
type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

// BaseClaims 基础用户信息
type BaseClaims struct {
	ID       int    `json:"id"`
	UserName string `json:"UserName"`
}

// JWT 错误定义
var (
	ErrTokenExpired     = errors.New("token 已过期")
	ErrTokenNotValidYet = errors.New("token 尚未生效")
	ErrTokenMalformed   = errors.New("token 格式错误")
	ErrTokenInvalid     = errors.New("token 无效")
	ErrTokenNotFound    = errors.New("token 未提供")
	ErrRefreshInvalid   = errors.New("刷新 token 不合法")
)

// jwtSigningKey
var jwtSigningKey []byte

// GetSigningKey 获取 JWT 签名密钥，使用单例模式确保只加载一次
func GetSigningKey() []byte {
	if jwtSigningKey == nil {
		jwtSigningKey = []byte(viper.GetString("Jwt.signingKey"))
	}
	return jwtSigningKey
}

// GetUserID 从 Context 中获取当前用户 ID
func GetUserID(ctx *echo.Context) (int, bool) {
	id, ok := ctx.Get("userID").(int)
	return id, ok
}

// GetUserName 从 Context 中获取当前用户名
func GetUserName(ctx *echo.Context) (string, bool) {
	name, ok := ctx.Get("UserName").(string)
	return name, ok
}

// GetClaims 从 Context 中获取完整 Claims（
func GetClaims(ctx *echo.Context) (*BaseClaims, bool) {
	claims, ok := ctx.Get("claims").(*BaseClaims)
	return claims, ok
}

func GetClaimsFromContext(ctx *echo.Context) (*BaseClaims, error) {
	tokenStr, err := ExtractToken(ctx)
	if err != nil {
		return nil, err
	}
	return ParseToken(tokenStr)
}

// ExtractToken 从请求头中提取 token
func ExtractToken(ctx *echo.Context) (string, error) {
	auth := ctx.Request().Header.Get("Authorization")
	if auth == "" {
		return "", ErrTokenNotFound
	}
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer "), nil
	}
	return auth, nil
}

// CreateClaims 构建 JWT Claims
func CreateClaims(base BaseClaims, expireHours int) CustomClaims {
	now := time.Now()
	return CustomClaims{
		BaseClaims: base,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * time.Duration(expireHours))),
			Issuer:    viper.GetString("Jwt.issuer"),
			Audience:  jwt.ClaimStrings{"Komu"},
		},
	}
}

// CreateToken 创建访问 token
func CreateToken(claims BaseClaims) (string, error) {
	c := CreateClaims(claims, viper.GetInt("Jwt.token_exp"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(GetSigningKey())
}

// CreateRefreshToken 创建刷新 token
func CreateRefreshToken(ctx context.Context, claims BaseClaims) (string, error) {
	redisKey := fmt.Sprintf("jwt:refresh:%d", claims.ID)
	oldToken, err := global.RedisClient.Get(ctx, redisKey).Result()
	if err == nil && oldToken != "" {
		if err := global.RedisClient.Del(ctx, redisKey).Err(); err != nil {
			return "", fmt.Errorf("删除旧的 refresh token 失败: %w", err)
		}
	}
	c := CreateClaims(claims, viper.GetInt("Jwt.refresh_exp"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	refreshToken, err := token.SignedString(GetSigningKey())
	if err != nil {
		return "", fmt.Errorf("生成 refresh token 失败: %w", err)
	}
	exp := time.Duration(viper.GetInt("Jwt.refresh_exp")) * time.Hour
	if err := global.RedisClient.Set(ctx, redisKey, refreshToken, exp).Err(); err != nil {
		return "", fmt.Errorf("存储 refresh token 到 Redis 失败: %w", err)
	}
	return refreshToken, nil
}

// GenerateTokenPair 生成 accessToken + refreshToken
func GenerateTokenPair(claims BaseClaims) (refreshToken, accessToken string, err error) {
	basicCtx := context.Background()
	refreshToken, err = CreateRefreshToken(basicCtx, claims)
	accessToken, err = CreateToken(claims)
	if err != nil {
		return "", "", fmt.Errorf("生成访问 token 失败: %w", err)
	}
	return refreshToken, accessToken, nil
}

// ParseToken 解析并验证 token，返回用户基础信息
func ParseToken(tokenString string) (*BaseClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方式: %v", token.Header["alg"])
		}
		return GetSigningKey(), nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, ErrTokenMalformed
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, ErrTokenExpired
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, ErrTokenNotValidYet
		default:
			return nil, ErrTokenInvalid
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return &claims.BaseClaims, nil
	}
	return nil, ErrTokenInvalid
}

// RefreshToken 使用 refreshToken 换取新的 accessToken
func RefreshToken(refreshToken string) (string, error) {
	claims, err := ParseToken(refreshToken)
	if err != nil {
		return "", err
	}
	redisKey := fmt.Sprintf("jwt:fresh:%d", claims.ID)
	stored, err := global.RedisClient.Get(context.Background(), redisKey).Result()
	if err != nil {
		return "", fmt.Errorf("查询刷新 token 失败: %w", err)
	}
	if stored != refreshToken {
		return "", ErrRefreshInvalid
	}
	return CreateToken(*claims)
}
