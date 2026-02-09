package router

import (
	"net/http"

	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initUserRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private *echo.Group) {
			public.POST("/login", api.UserApi.Login)
			private.GET("/hello", func(c *echo.Context) error {
				return c.JSON(http.StatusOK, "jello")
			})
		},
	)
}
