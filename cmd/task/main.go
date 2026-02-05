package main

import (
	"context"

	"github.com/sllt/kite-layout/cmd/task/wire"
	"github.com/sllt/kite-layout/pkg/config"
)

func main() {
	// Ensure Kite can find the .env config file
	config.SetupKiteEnv()

	app, cleanup, err := wire.NewWire()
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
