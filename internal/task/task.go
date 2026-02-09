package task

import (
	"github.com/sllt/kite-layout/internal/repository"
	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite-layout/pkg/sid"
)

type Task struct {
	logger *log.Logger
	sid    *sid.Sid
	tm     repository.Transaction
}

func NewTask(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Task {
	return &Task{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
