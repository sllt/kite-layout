package server

import (
	"github.com/sllt/kite-layout/internal/middleware"
	"github.com/sllt/kite-layout/internal/router"
	"github.com/sllt/kite/pkg/kite"
)

// HTTPServerReady is a marker type indicating the HTTP server has been configured.
// It is used in the Wire DI graph to ensure route and middleware registration
// happens before the application starts.
type HTTPServerReady struct{}

func NewHTTPServer(
	deps router.RouterDeps,
) *HTTPServerReady {
	app := deps.App

	// Register global middleware
	app.UseMiddleware(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(deps.Logger),
		middleware.RequestLogMiddleware(deps.Logger),
		// NoStrictAuth runs globally: extracts JWT claims into context when token is present,
		// but does not reject requests without a token.
		middleware.NoStrictAuth(deps.JWT, deps.Logger),
	)

	// Root route
	app.GET("/", func(ctx *kite.Context) (any, error) {
		return map[string]interface{}{
			":)": "Thank you for using nunu!",
		}, nil
	})

	// Register user routes
	router.InitUserRouter(deps)

	return &HTTPServerReady{}
}
