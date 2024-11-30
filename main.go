package main

import (
	// "time"

	"context"
	"net/http"
	"os"
	"os/signal"

	"movie-festival-app/entity"
	appInit "movie-festival-app/init"
	"movie-festival-app/module/handler"
	"movie-festival-app/module/store"
	"movie-festival-app/module/usecase"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	config *entity.Config
)

func init() {
	appInit.SetupLogger()
	config = appInit.SetupMainConfig()
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover(), middleware.Logger(), middleware.RequestID(), middleware.Secure())
	main := e.Group("/movie-festival-app")

	handler.New(main, &usecase.Methods{
		Stores: store.New(config),
		Config: config,
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(config.API.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), config.Context.Timeout*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
