//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/sllt/kite-layout/internal/server"
	"github.com/sllt/kite-layout/pkg/app"
	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite/pkg/kite"
)

var serverSet = wire.NewSet(
	server.NewMigrateServer,
)

// NewKiteApp creates a kite.App.
func NewKiteApp() *kite.App {
	return kite.New()
}

// NewLogger extracts kite's logger from the container and wraps it.
func NewLogger(app *kite.App) *log.Logger {
	return log.NewLogger(app.Container().Logger)
}

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
		NewKiteApp,
		NewLogger,
		newApp,
	))
}
