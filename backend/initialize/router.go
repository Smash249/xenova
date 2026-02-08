package initialize

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func initRouter() {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	public := e.Group("/public")
	private := e.Group("/private")
	fmt.Print(public, private)
}
