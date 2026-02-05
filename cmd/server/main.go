package main

import (
	"github.com/sllt/kite-layout/cmd/server/wire"
	"github.com/sllt/kite-layout/pkg/config"
)

// @title           Nunu Example API
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8000
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Ensure Kite can find the .env config file
	config.SetupKiteEnv()

	app, cleanup, err := wire.NewWire()
	defer cleanup()
	if err != nil {
		panic(err)
	}

	// Kite's Run() manages the HTTP server lifecycle including graceful shutdown
	app.Run()
}
