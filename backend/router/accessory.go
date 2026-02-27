package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initAccessoryRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private, admin *echo.Group) {
			public.GET("/accessory_series", api.AccessoryApi.GetAccessorySeries)
			admin.POST("/accessory_series", api.AccessoryApi.CreateAccessorySeries)
			admin.PUT("/accessory_series", api.AccessoryApi.UpdateAccessorySeries)
			admin.DELETE("/accessory_series", api.AccessoryApi.DeleteAccessorySeries)

			public.GET("/accessorys", api.AccessoryApi.GetAccessoryList)
			private.GET("/accessorys/:accessory_id", api.AccessoryApi.GetAccessoryDetail)
			admin.POST("/accessorys", api.AccessoryApi.CreateAccessory)
			admin.PUT("/accessorys", api.AccessoryApi.UpdateAccessory)
			admin.DELETE("/accessorys", api.AccessoryApi.DeleteAccessory)
		},
	)
}
