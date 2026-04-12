//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/sllt/kite-layout/internal/bootstrap"
	"github.com/sllt/kite-layout/internal/handler"
	"github.com/sllt/kite-layout/internal/repository"
	"github.com/sllt/kite-layout/internal/router"
	"github.com/sllt/kite-layout/internal/server"
	"github.com/sllt/kite-layout/internal/service"
	"github.com/sllt/kite-layout/pkg/jwt"
	"github.com/sllt/kite-layout/pkg/sid"
	"github.com/google/wire"
	"github.com/sllt/kite/pkg/kite"
)

var repositorySet = wire.NewSet(
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
)

// App wraps kite.App, embedding it so callers can use it transparently.
// This wrapper exists because wire requires distinct types for providers.
type App struct {
	*kite.App
}

func newApp(app *kite.App, _ *server.HTTPServerReady) *App {
	return &App{App: app}
}

func NewWire() (*App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		bootstrap.NewKiteApp,
		bootstrap.NewLogger,
		bootstrap.NewDB,
		wire.Struct(new(router.RouterDeps), "*"),
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
