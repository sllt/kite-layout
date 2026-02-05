package http

import (
	"context"
	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite/pkg/kite"
)

// Server wraps kite.App for compatibility with existing server management code
type Server struct {
	*kite.App
	logger *log.Logger
}

type Option func(s *Server)

// NewServer creates a new HTTP server wrapping kite.App
func NewServer(app *kite.App, logger *log.Logger, opts ...Option) *Server {
	s := &Server{
		App:    app,
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// WithServerHost sets the server host (configured via HTTP_HOST env var in Kite)
func WithServerHost(host string) Option {
	return func(s *Server) {
		// Kite uses HTTP_HOST environment variable for configuration
		// This option is kept for backward compatibility but may not be used directly
	}
}

// WithServerPort sets the server port (configured via HTTP_PORT env var in Kite)
func WithServerPort(port int) Option {
	return func(s *Server) {
		// Kite uses HTTP_PORT environment variable for configuration
		// This option is kept for backward compatibility but may not be used directly
	}
}

// Start starts the HTTP server
func (s *Server) Start(ctx context.Context) error {
	s.App.Run()
	return nil
}

// Stop gracefully stops the HTTP server
func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("Shutting down server...")
	// Kite handles graceful shutdown internally
	s.logger.Info("Server exiting")
	return nil
}
