package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initSoftwareRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private *echo.Group) {
			public.GET("/software_series", api.SoftwareApi.GetSoftwareSeries)
			private.POST("/software_series", api.SoftwareApi.CreateSoftwareSeries)
			private.PUT("/software_series", api.SoftwareApi.UpdateSoftwareSeries)
			private.DELETE("/software_series", api.SoftwareApi.DeleteSoftwareSeries)

			public.GET("/softwares", api.SoftwareApi.GetSoftwareList)
			public.GET("/softwares/:software_id", api.SoftwareApi.GetSoftwareDetail)
			private.POST("/softwares", api.SoftwareApi.CreateSoftware)
			private.PUT("/softwares", api.SoftwareApi.UpdateSoftware)
			private.DELETE("/softwares", api.SoftwareApi.DeleteSoftware)
		},
	)
}
