package main

import (
	"github.com/Leaeraso/max_inventory/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// dentro de provide se pasan todas las funciones que devuelvan un struct
		fx.Provide(
			settings.New,
		),
		// dentro de invoke se ejecutan los comandos antes de que la aplicacion empiece a correr
		fx.Invoke(),
	)

	app.Run()
}
