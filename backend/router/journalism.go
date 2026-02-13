package router

import (
	"github.com/Smash249/xenova/backend/internal/api"
	"github.com/labstack/echo/v5"
)

func initJournalismRouter() {
	GroupRouterHubApp.RegisterRouterHub(
		func(public, private *echo.Group) {
			public.GET("/journalism_series", api.JournalismApi.GetJournalismSeries)
			private.POST("/journalism_series", api.JournalismApi.CreateJournalismSeries)
			private.PUT("/journalism_series", api.JournalismApi.UpdateJournalismSeries)
			private.DELETE("/journalism_series", api.JournalismApi.DeleteJournalismSeries)

			public.GET("/journalisms", api.JournalismApi.GetJournalismList)
			public.GET("/journalisms/:journalism_id", api.JournalismApi.GetJournalismDetail)
			private.POST("/journalisms", api.JournalismApi.CreateJournalism)
			private.PUT("/journalisms", api.JournalismApi.UpdateJournalism)
			private.DELETE("/journalisms", api.JournalismApi.DeleteJournalism)
		},
	)
}
