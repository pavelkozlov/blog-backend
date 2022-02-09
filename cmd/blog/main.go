package main

import (
	"blog/internal/post/delivery"
	"blog/pkg/database"
	"context"
	"net/http"

	"blog/pkg/logger"

	"os"

	"os/signal"

	"sync"

	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"blog/pkg/config"

	"syscall"
)

func main() {

	cfg := config.NewConfig()
	log := logger.NewLogger(cfg)
	log.Debug("config loaded", "lol", "kek", "cfg", *cfg)
	db := database.NewPostgresDB(cfg)
	if db != nil {
		log.Debug("Database connected")
	}

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
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/blog", delivery.AllPosts)

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
