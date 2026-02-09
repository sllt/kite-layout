package app

import (
	"context"
	"github.com/sllt/kite-layout/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	name    string
	servers []server.Server
}

type Option func(a *App)

func NewApp(opts ...Option) *App {
	a := &App{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func WithServer(servers ...server.Server) Option {
	return func(a *App) {
		a.servers = servers
	}
}

func WithName(name string) Option {
	return func(a *App) {
		a.name = name
	}
}

func (a *App) Run(ctx context.Context) error {
	if len(a.servers) == 0 {
		return nil
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signals)

	startErrCh := make(chan error, len(a.servers))

	for _, srv := range a.servers {
		go func(srv server.Server) {
			startErrCh <- srv.Start(ctx)
		}(srv)
	}

	var runErr error

	select {
	case <-signals:
		// Received termination signal
		log.Println("Received termination signal")
		cancel()
	case <-ctx.Done():
		// Context canceled
		log.Println("Context canceled")
	case err := <-startErrCh:
		// One-shot server (e.g. migration) finished, or server startup failed.
		if err != nil {
			runErr = err
			log.Printf("Server start err: %v", err)
		}
		cancel()
	}

	// Gracefully stop the servers
	for _, srv := range a.servers {
		err := srv.Stop(ctx)
		if err != nil {
			log.Printf("Server stop err: %v", err)
			if runErr == nil {
				runErr = err
			}
		}
	}

	return runErr
}
