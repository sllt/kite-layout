package repository

import (
	"context"
	"database/sql"

	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite/pkg/kite/infra"
	kiteSQL "github.com/sllt/kite/pkg/kite/datasource/sql"
)

const ctxTxKey = "TxKey"

// Querier is a common interface satisfied by both infra.DB and kiteSQL.Tx
type Querier interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Repository struct {
	db     infra.DB
	logger *log.Logger
}

func NewRepository(logger *log.Logger, db infra.DB) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(r *Repository) Transaction {
	return r
}

// GetQuerier returns the transaction from context if available, otherwise the DB.
func (r *Repository) GetQuerier(ctx context.Context) Querier {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*kiteSQL.Tx); ok {
			return tx
		}
	}
	return r.db
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	txCtx := context.WithValue(ctx, ctxTxKey, tx)
	if err := fn(txCtx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
