package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initSoftwareRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private *echo.Group) {
			public.GET("/software_categories", api.SoftwareApi.GetSoftwareCategory)
			private.POST("/software_categories", api.SoftwareApi.CreateSoftwareCategory)
			private.PUT("/software_categories", api.SoftwareApi.UpdateSoftwareCategory)
			private.DELETE("/software_categories", api.SoftwareApi.DeleteSoftwareCategory)

			public.GET("/softwares", api.SoftwareApi.GetSoftwareList)
			public.GET("/softwares/:software_id", api.SoftwareApi.GetSoftwareDetail)
			private.POST("/softwares", api.SoftwareApi.CreateSoftware)
			private.PUT("/softwares", api.SoftwareApi.UpdateSoftware)
			private.DELETE("/softwares", api.SoftwareApi.DeleteSoftware)
		},
	)
}
