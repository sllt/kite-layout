//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/sllt/kite-layout/internal/bootstrap"
	"github.com/sllt/kite-layout/internal/repository"
	"github.com/sllt/kite-layout/internal/server"
	"github.com/sllt/kite-layout/internal/task"
	"github.com/sllt/kite-layout/pkg/app"
	"github.com/sllt/kite-layout/pkg/sid"
	"github.com/google/wire"
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
		bootstrap.NewKiteApp,
		bootstrap.NewLogger,
		bootstrap.NewDB,
		newApp,
		sid.NewSid,
	))
}
