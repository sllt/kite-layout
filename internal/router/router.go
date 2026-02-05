package router

import (
	"github.com/sllt/kite-layout/internal/handler"
	"github.com/sllt/kite-layout/pkg/jwt"
	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite/pkg/kite"
)

type RouterDeps struct {
	App         *kite.App
	Logger      *log.Logger
	JWT         *jwt.JWT
	UserHandler *handler.UserHandler
}
