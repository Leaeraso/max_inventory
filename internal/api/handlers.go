package api

import (
	"net/http"

	"github.com/Leaeraso/max_inventory/internal/api/dtos"
	"github.com/Leaeraso/max_inventory/internal/service"
	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}
	// recupera los datos del cuerpo de la request y los pasa a params
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	err = a.dateValidator.Struct(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "user already exists"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}
	
	return c.JSON(http.StatusCreated, nil)
}

