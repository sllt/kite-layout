package router

// InitUserRouter registers user routes on the Kite app.
// Since Kite doesn't support route-level middleware groups, authentication
// is handled globally via NoStrictAuth middleware (which sets claims when
// a token is present but doesn't reject requests without one).
// Routes requiring strict authentication check for claims in the handler.
func InitUserRouter(deps RouterDeps) {
	app := deps.App

	// No authentication required
	app.POST("/api/v1/register", deps.UserHandler.Register)
	app.POST("/api/v1/login", deps.UserHandler.Login)

	// Token optional - NoStrictAuth middleware runs globally and sets claims if present.
	// Handler checks for claims and returns ErrUnauthorized if missing.
	app.GET("/api/v1/user", deps.UserHandler.GetProfile)

	// Token required - NoStrictAuth middleware runs globally and sets claims if present.
	// Handler checks for claims and returns ErrUnauthorized if missing.
	app.PUT("/api/v1/user", deps.UserHandler.UpdateProfile)
}
