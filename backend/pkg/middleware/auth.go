package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Smash249/xenova/backend/utils"
	"github.com/labstack/echo/v5"
)

type AuthConfig struct {
	Skipper     func(ctx *echo.Context) bool
	WhiteList   []string
	TokenLookup string
}

var DefaultAuthConfig = AuthConfig{
	TokenLookup: "header:Authorization",
}

// NewAuth 返回一个 JWT 认证中间件，默认从 Authorization header 中提取 token
func NewAuth() echo.MiddlewareFunc {
	return NewAuthWithConfig(DefaultAuthConfig)
}

func NewAuthWithConfig(config AuthConfig) echo.MiddlewareFunc {
	if config.TokenLookup == "" {
		config.TokenLookup = DefaultAuthConfig.TokenLookup
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx *echo.Context) error {
			// 跳过检查
			if config.Skipper != nil && config.Skipper(ctx) {
				return next(ctx)
			}

			// 提取 token
			tokenStr, err := extractTokenFromLookup(ctx, config.TokenLookup)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "缺少认证信息")
			}

			// 解析验证 token
			claims, err := utils.ParseToken(tokenStr)
			if err != nil {
				switch {
				case errors.Is(err, utils.ErrTokenExpired):
					return echo.NewHTTPError(http.StatusUnauthorized, "token 已过期，请重新登录")
				case errors.Is(err, utils.ErrTokenMalformed):
					return echo.NewHTTPError(http.StatusUnauthorized, "token 格式错误")
				case errors.Is(err, utils.ErrTokenNotValidYet):
					return echo.NewHTTPError(http.StatusUnauthorized, "token 尚未生效")
				default:
					return echo.NewHTTPError(http.StatusUnauthorized, "token 无效")
				}
			}
			ctx.Set("userID", claims.ID)
			ctx.Set("userName", claims.UserName)
			ctx.Set("role", claims.Role)
			ctx.Set("claims", claims)
			return next(ctx)
		}
	}
}

// extractTokenFromLookup
// 格式: "header:Authorization", "query:token", "cookie:token"
func extractTokenFromLookup(ctx *echo.Context, lookup string) (string, error) {
	parts := strings.SplitN(lookup, ":", 2)
	if len(parts) != 2 {
		return "", utils.ErrTokenNotFound
	}
	switch parts[0] {
	case "header":
		return utils.ExtractToken(ctx)
	case "query":
		token := ctx.QueryParam(parts[1])
		if token == "" {
			return "", utils.ErrTokenNotFound
		}
		return token, nil
	case "cookie":
		cookie, err := ctx.Cookie(parts[1])
		if err != nil || cookie.Value == "" {
			return "", utils.ErrTokenNotFound
		}
		return cookie.Value, nil
	default:
		return "", utils.ErrTokenNotFound
	}
}
