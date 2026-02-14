package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initUserRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private *echo.Group) {
			public.POST("/login", api.UserApi.Login)
			public.POST("/register", api.UserApi.Register)
			public.GET("/health", api.UserApi.Health)
		},
	)
}
