package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initJournalism() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private, admin *echo.Group) {
			public.GET("/journalism_series", api.JournalismApi.GetJournalismSeries)
			admin.POST("/journalism_series", api.JournalismApi.CreateJournalismSeries)
			admin.PUT("/journalism_series", api.JournalismApi.UpdateJournalismSeries)
			admin.DELETE("/journalism_series", api.JournalismApi.DeleteJournalismSeries)

			public.GET("/journalisms", api.JournalismApi.GetJournalismList)
			private.GET("/journalisms/:journalism_id", api.JournalismApi.GetJournalismDetail)
			admin.POST("/journalisms", api.JournalismApi.CreateJournalism)
			admin.PUT("/journalisms", api.JournalismApi.UpdateJournalism)
			admin.DELETE("/journalisms", api.JournalismApi.DeleteJournalism)
		})
}
