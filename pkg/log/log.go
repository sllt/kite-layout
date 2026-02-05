package log

import (
	"github.com/sllt/kite/pkg/kite/logging"
)

// Logger wraps kite's logging.Logger so existing code can keep using *log.Logger
// without changing function signatures.
type Logger struct {
	logging.Logger
}

// NewLogger creates a Logger backed by the given kite logging.Logger.
func NewLogger(l logging.Logger) *Logger {
	return &Logger{Logger: l}
}
