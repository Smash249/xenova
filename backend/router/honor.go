package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initHonorRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private *echo.Group) {
			public.GET("/company_honor", api.HonorApi.GetCompanyHonorList)
			private.POST("/company_honor", api.HonorApi.CreateCompanyHonor)
			private.PUT("/company_honor", api.HonorApi.UpdateCompanyHonor)
			private.DELETE("/company_honor", api.HonorApi.DeleteCompanyHonor)

			public.GET("/love_activity", api.HonorApi.GetLoveActivityList)
			private.POST("/love_activity", api.HonorApi.CreateLoveActivity)
			private.PUT("/love_activity", api.HonorApi.UpdateLoveActivity)
			private.DELETE("/love_activity", api.HonorApi.DeleteLoveActivity)

			public.GET("/company_patnet", api.HonorApi.GetCompanyPatentList)
			private.POST("/company_patent", api.HonorApi.CreateCompanyPatent)
			private.PUT("/company_patent", api.HonorApi.UpdateCompanyPatent)
			private.DELETE("/company_patent", api.HonorApi.DeleteCompanyPatent)

		})
}
