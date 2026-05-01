package server

import (
	"github.com/sllt/kite-layout/internal/grpc/user"
	"github.com/sllt/kite-layout/internal/middleware"
	"github.com/sllt/kite-layout/internal/router"
	"github.com/sllt/kite-layout/internal/service"
	"github.com/sllt/kite/pkg/kite"
)

func NewHTTPServer(
	deps router.RouterDeps,
	userService service.UserService,
) {
	app := deps.App

	// Register global middleware
	app.Use(
		middleware.CORSMiddleware(),
		// NoStrictAuth runs globally: extracts JWT claims into context when token is present,
		// but does not reject requests without a token. Strict routes are enforced by
		// group-scoped Kite middleware in internal/router.
		middleware.NoStrictAuth(deps.JWT, deps.Logger),
	)

	// Root route
	app.GET("/", func(ctx *kite.Context) (any, error) {
		return map[string]any{
			":)": "Thank you for using kite!",
		}, nil
	})

	// Register user routes (HTTP)
	router.InitUserRouter(deps)

	// Register gRPC services
	userKiteServer := user.NewUserServiceKiteServerWithService(userService)
	user.RegisterUserServiceServerWithKite(app, userKiteServer)
}
