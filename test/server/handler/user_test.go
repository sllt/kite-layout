package handler

// TODO: Handler tests need kite testing infrastructure.
// kite.Context cannot be created externally, so handler tests
// require either:
// 1. A TestContext helper in the kite package
// 2. Starting a real kite server via httptest
// 3. Testing through integration tests
//
// The handler signatures are now: func(ctx *kite.Context) (any, error)
// which can only be invoked through kite's internal routing.
