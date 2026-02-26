package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initSystemRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private, admin *echo.Group) {
			public.GET("/health", api.SystemApi.Health)
			admin.GET("/users", api.SystemApi.GetSystemUserList)
			admin.GET("/users/:id", api.SystemApi.BanedUserById)

			admin.GET("/messages", api.SystemApi.GetSystemMessageList)
			private.POST("/messages", api.SystemApi.CreateSystemMessage)
			admin.PUT("/messages", api.SystemApi.UpdateSystemMessage)
			admin.DELETE("/messages", api.SystemApi.DeleteSystemMessage)

			public.GET("/job_positions", api.SystemApi.GetSystemJobPositionList)
			admin.POST("/job_positions", api.SystemApi.CreateSystemJobPosition)
			admin.PUT("/job_positions", api.SystemApi.UpdateSystemJobPosition)
			admin.DELETE("/job_positions", api.SystemApi.DeleteSystemJobPosition)
		},
	)
}
