package bootstrap

import (
	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite/pkg/kite"
	"github.com/sllt/kite/pkg/kite/infra"
)

// NewKiteApp creates a new kite.App.
func NewKiteApp() *kite.App {
	return kite.New()
}

// NewLogger extracts kite's logger from the container and wraps it.
func NewLogger(app *kite.App) *log.Logger {
	return log.NewLogger(app.Container().Logger)
}

// NewDB extracts infra.DB from the kite app's container.
func NewDB(app *kite.App) infra.DB {
	return app.Container().SQL
}
