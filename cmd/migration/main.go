package main

import (
	"github.com/sllt/kite-layout/internal/bootstrap"
	"github.com/sllt/kite-layout/internal/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bootstrap.CoreModule,
		fx.Invoke(server.RegisterMigrateServer),
	).Run()
}
