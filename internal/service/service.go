package service

import (
	"github.com/sllt/kite-layout/internal/repository"
	"github.com/sllt/kite-layout/pkg/jwt"
	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite-layout/pkg/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
