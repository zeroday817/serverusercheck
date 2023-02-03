package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	svc Service
}

func New(svc Service) *Endpoint {
	return &Endpoint{
		svc: svc,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {

	status := fmt.Sprintf("Days left until 25 : %d", e.svc.DaysLeft())

	err := ctx.String(http.StatusOK, status)
	if err != nil {
		return err
	}
	return nil
}
