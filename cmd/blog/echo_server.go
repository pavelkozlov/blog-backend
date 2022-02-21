package main

import (
	"blog/cmd/di"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initEchoServer(app *di.Application) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `{"time"="${time_rfc3339}", "method"="${method}", "path"="${path}", "id"="${id}", "status"="${status}"}'` + "\n",
		CustomTimeFormat: "",
		Output:           os.Stderr,
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))

	e.GET("/blog", echo.WrapHandler(app.BlogHandlers.AllPosts()))

	return e
}
