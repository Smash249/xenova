package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initSoftwareRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private, admin *echo.Group) {
			public.GET("/software_series", api.SoftwareApi.GetSoftwareSeries)
			admin.POST("/software_series", api.SoftwareApi.CreateSoftwareSeries)
			admin.PUT("/software_series", api.SoftwareApi.UpdateSoftwareSeries)
			admin.DELETE("/software_series", api.SoftwareApi.DeleteSoftwareSeries)

			public.GET("/softwares", api.SoftwareApi.GetSoftwareList)
			private.GET("/softwares/:software_id", api.SoftwareApi.GetSoftwareDetail)
			admin.POST("/softwares", api.SoftwareApi.CreateSoftware)
			admin.PUT("/softwares", api.SoftwareApi.UpdateSoftware)
			admin.DELETE("/softwares", api.SoftwareApi.DeleteSoftware)
		},
	)
}
