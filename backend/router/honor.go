package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initHonorRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private, admin *echo.Group) {
			public.GET("/company_honor", api.HonorApi.GetCompanyHonorList)
			admin.POST("/company_honor", api.HonorApi.CreateCompanyHonor)
			admin.PUT("/company_honor", api.HonorApi.UpdateCompanyHonor)
			admin.DELETE("/company_honor", api.HonorApi.DeleteCompanyHonor)

			public.GET("/love_activity", api.HonorApi.GetLoveActivityList)
			admin.POST("/love_activity", api.HonorApi.CreateLoveActivity)
			admin.PUT("/love_activity", api.HonorApi.UpdateLoveActivity)
			admin.DELETE("/love_activity", api.HonorApi.DeleteLoveActivity)

			public.GET("/company_patnet", api.HonorApi.GetCompanyPatentList)
			admin.POST("/company_patent", api.HonorApi.CreateCompanyPatent)
			admin.PUT("/company_patent", api.HonorApi.UpdateCompanyPatent)
			admin.DELETE("/company_patent", api.HonorApi.DeleteCompanyPatent)

		})
}
