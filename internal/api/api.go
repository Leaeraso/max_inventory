package api

import (
	"github.com/Leaeraso/max_inventory/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type API struct {
	serv service.Service
	dateValidator *validator.Validate
}

func New(s service.Service) *API {
	return &API{
		serv: s,
		dateValidator: validator.New(),
	}
}

// reciever de nuestro API
func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)
	return e.Start(address)
}