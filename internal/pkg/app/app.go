package app

import (
	"GOLANG_PROJECTS/internal/app/endpoint"
	"GOLANG_PROJECTS/internal/app/mw"
	"GOLANG_PROJECTS/internal/app/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	svc  *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	app := &App{}
	app.svc = service.New()
	app.e = endpoint.New(app.svc)

	app.echo = echo.New()
	app.echo.Use(mw.RoleCheck)

	app.echo.GET("/status", app.e.Status)

	return app, nil
}

func (app *App) Run() error {
	fmt.Println("server running")

	err := app.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
