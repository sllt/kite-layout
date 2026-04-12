package main

import (
	"context"

	"github.com/sllt/kite-layout/cmd/task/wire"
)

func main() {
	app, cleanup, err := wire.NewWire()
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
