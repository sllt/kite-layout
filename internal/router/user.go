package router

import (
	"github.com/sllt/kite-layout/internal/handler"
	"github.com/sllt/kite-layout/pkg/errcode"
	"github.com/sllt/kite/pkg/kite"
)

func requireAuthMiddleware() kite.KiteMiddleware {
	return func(next kite.Handler) kite.Handler {
		return func(ctx *kite.Context) (any, error) {
			if handler.GetUserIdFromCtx(ctx) == "" {
				return nil, errcode.ErrUnauthorized
			}
			return next(ctx)
		}
	}
}

// InitUserRouter registers user routes on the Kite app.
// Token parsing is handled globally via NoStrictAuth middleware (which sets claims
// when a token is present but doesn't reject requests without one).
// Protected routes use a group-scoped Kite middleware for strict auth.
func InitUserRouter(deps RouterDeps) {
	apiV1 := deps.App.Group("/api/v1")

	// No authentication required
	apiV1.POST("/register", deps.UserHandler.Register)
	apiV1.POST("/login", deps.UserHandler.Login)

	// Token required
	userGroup := apiV1.Group("/user")
	userGroup.UseMiddleware(requireAuthMiddleware())
	userGroup.GET("/", deps.UserHandler.GetProfile)
	userGroup.PUT("/", deps.UserHandler.UpdateProfile)
}
