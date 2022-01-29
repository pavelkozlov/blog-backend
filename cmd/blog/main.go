package main

import (
	"blog/pkg/config"
	"context"
	"fmt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	cfg := config.NewConfig()
	fmt.Println(cfg)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `{"time"="${time_rfc3339}", "method"="${method}", "path"="${path}", "id"="${id}", "status"="${status}"}'` + "\n",
		CustomTimeFormat: "",
		Output:           os.Stderr,
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go shutdown(e, wg)

	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server with err: ", err.Error())
	}

	wg.Wait()
}

func shutdown(e *echo.Echo, wg *sync.WaitGroup) {
	defer wg.Done()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
