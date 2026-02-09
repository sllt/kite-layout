package server

import (
	"context"

	"github.com/sllt/kite-layout/migrations"
	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite/pkg/kite"
)

type MigrateServer struct {
	app *kite.App
	log *log.Logger
}

func NewMigrateServer(app *kite.App, log *log.Logger) *MigrateServer {
	return &MigrateServer{
		app: app,
		log: log,
	}
}
func (m *MigrateServer) Start(ctx context.Context) error {
	_ = ctx
	m.app.Migrate(migrations.All())
	m.log.Info("Migration success")
	return nil
}
func (m *MigrateServer) Stop(ctx context.Context) error {
	m.log.Info("Migration stop")
	return nil
}
