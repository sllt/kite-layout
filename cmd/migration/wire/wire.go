//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/sllt/kite-layout/internal/bootstrap"
	"github.com/sllt/kite-layout/internal/server"
	"github.com/sllt/kite-layout/pkg/app"
)

var serverSet = wire.NewSet(
	server.NewMigrateServer,
)

// build App
func newApp(
	migrateServer *server.MigrateServer,
) *app.App {
	return app.NewApp(
		app.WithServer(migrateServer),
		app.WithName("demo-migrate"),
	)
}

func NewWire() (*app.App, func(), error) {
	panic(wire.Build(
		serverSet,
		bootstrap.NewKiteApp,
		bootstrap.NewLogger,
		newApp,
	))
}
