package main

import (
	"blog/cmd/di"
	"context"
	"log"
	"net/http"

	"os"

	"os/signal"

	"sync"

	"time"

	"syscall"

	echo "github.com/labstack/echo/v4"
)

func main() {

	app, err := di.NewContainer()
	if err != nil {
		log.Fatalf("Can not init application container: %v\n", err)
	}

	e := initEchoServer(app)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go shutdown(e, wg)
	if err = e.Start(":8080"); err != nil && err != http.ErrServerClosed {
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
