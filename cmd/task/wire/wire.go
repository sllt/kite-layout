//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/sllt/kite-layout/internal/repository"
	"github.com/sllt/kite-layout/internal/server"
	"github.com/sllt/kite-layout/internal/task"
	"github.com/sllt/kite-layout/pkg/app"
	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite-layout/pkg/sid"
	"github.com/google/wire"
	"github.com/sllt/kite/pkg/kite"
	"github.com/sllt/kite/pkg/kite/infra"
)

var repositorySet = wire.NewSet(
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var taskSet = wire.NewSet(
	task.NewTask,
	task.NewUserTask,
)
var serverSet = wire.NewSet(
	server.NewTaskServer,
)

// NewKiteApp creates a kite.App.
func NewKiteApp() *kite.App {
	return kite.New()
}

// NewLogger extracts kite's logger from the container and wraps it.
func NewLogger(app *kite.App) *log.Logger {
	return log.NewLogger(app.Container().Logger)
}

// NewDB extracts infra.DB from the kite app's container
func NewDB(app *kite.App) infra.DB {
	return app.Container().SQL
}

// build App
func newApp(
	task *server.TaskServer,
) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("demo-task"),
	)
}

func NewWire() (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		taskSet,
		serverSet,
		NewKiteApp,
		NewLogger,
		NewDB,
		newApp,
		sid.NewSid,
	))
}
