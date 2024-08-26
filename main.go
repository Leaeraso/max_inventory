package main

import (
	"context"

	"github.com/Leaeraso/max_inventory/database"
	"github.com/Leaeraso/max_inventory/internal/repository"
	"github.com/Leaeraso/max_inventory/internal/service"
	"github.com/Leaeraso/max_inventory/settings"
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
		),
		// dentro de invoke se ejecutan los comandos antes de que la aplicacion empiece a correr
		fx.Invoke(
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(ctx, "my@email.com", "myname", "mypassword")
				if err != nil {
					panic(err)
				}

				u, err := serv.LoginUser(ctx, "my@email.com", "mypassword")
				if err != nil {
					panic(err)
				}

				if u.Name != "myname" {
					panic("wrong name")
				}
			},
		),
	)

	app.Run()
}
