package main

import (
	"context"

	"github.com/Leaeraso/max_inventory/database"
	"github.com/Leaeraso/max_inventory/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// dentro de provide se pasan todas las funciones que devuelvan un struct
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		// dentro de invoke se ejecutan los comandos antes de que la aplicacion empiece a correr
		fx.Invoke(
			func (db *sqlx.DB) {
				_, err := db.Query("select * from users")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
