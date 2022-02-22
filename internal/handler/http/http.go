package http

import (
	"context"
	netHTTP "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/mahbubzulkarnain/golang-singleflight-example/internal/service"
)

func RegisterHTTPServer(port string, s service.Service) {
	e := echo.New()

	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())

	v1 := e.Group("/v1")
	{
		Handler{UserService: s.V1.UserService}.Route(v1)
	}

	// Start server
	go func() {
		if err := e.Start(":" + port); err != nil && err != netHTTP.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	//Recieve shutdown signals.
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
