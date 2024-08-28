package main

import (
	"context"
	"fmt"

	"github.com/Leaeraso/max_inventory/database"
	"github.com/Leaeraso/max_inventory/internal/api"
	"github.com/Leaeraso/max_inventory/internal/repository"
	"github.com/Leaeraso/max_inventory/internal/service"
	"github.com/Leaeraso/max_inventory/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// dentro de provide se pasan todas las funciones que devuelvan un struct
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		// dentro de invoke se ejecutan los comandos antes de que la aplicacion empiece a correr
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

// configuramos el ciclo de vida de la app fx
func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
