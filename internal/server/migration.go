package server

import (
	"context"
	"os"

	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite/pkg/kite/infra"
)

type MigrateServer struct {
	db  infra.DB
	log *log.Logger
}

func NewMigrateServer(db infra.DB, log *log.Logger) *MigrateServer {
	return &MigrateServer{
		db:  db,
		log: log,
	}
}
func (m *MigrateServer) Start(ctx context.Context) error {
	_, err := m.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id TEXT UNIQUE NOT NULL,
			nickname TEXT NOT NULL DEFAULT '',
			password TEXT NOT NULL,
			email TEXT NOT NULL,
			created_at DATETIME,
			updated_at DATETIME
		)
	`)
	if err != nil {
		m.log.Errorf("migrate error: %v", err)
		return err
	}
	m.log.Info("Migration success")
	os.Exit(0)
	return nil
}
func (m *MigrateServer) Stop(ctx context.Context) error {
	m.log.Info("Migration stop")
	return nil
}
