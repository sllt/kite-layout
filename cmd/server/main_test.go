package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

type stubFXApp struct {
	stopErr error
	stopCtx context.Context
}

func (s *stubFXApp) Stop(ctx context.Context) error {
	s.stopCtx = ctx
	return s.stopErr
}

func TestStopFXAppReturnsStopError(t *testing.T) {
	expectedErr := errors.New("stop failed")
	app := &stubFXApp{stopErr: expectedErr}

	err := stopFXApp(app, time.Millisecond)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected stop error %v, got %v", expectedErr, err)
	}
}

func TestStopFXAppUsesTimeoutContext(t *testing.T) {
	app := &stubFXApp{}

	if err := stopFXApp(app, time.Second); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if app.stopCtx == nil {
		t.Fatal("expected stop context to be passed")
	}
	if _, ok := app.stopCtx.Deadline(); !ok {
		t.Fatal("expected stop context to have a deadline")
	}
}
