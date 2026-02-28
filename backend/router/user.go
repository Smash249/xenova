package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initUserRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private, admin *echo.Group) {
			public.POST("/login", api.UserApi.FrontendLogin)
			public.POST("/code", api.UserApi.SendEmailCode)
			public.POST("/register", api.UserApi.Register)
			public.POST("/sigin", api.UserApi.AdminLogin)

			private.GET("/userInfo", api.UserApi.GetUserInfo)
			private.PUT("/userInfo", api.UserApi.UpdateUserInfo)
			admin.POST("/userInfo", api.UserApi.AdminUpdateUserInfo)
			private.POST("/changePassword", api.UserApi.ChangePassword)
			private.POST("/upload", api.UserApi.Upload)
		},
	)
}
