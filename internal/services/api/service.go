package api

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service
type Service interface {
	Auth(ctx context.Context, param1, param2 string) ([]string, error)
	Credit(ctx context.Context, param1, param2 string) ([]string, error)
	Debit(ctx context.Context, id, param1, param2 string) ([]string, error)
}

func New(uuidFunc func() uuid.UUID, timeFunc func() time.Time) Service {
	return &service{
		uuidFunc: uuidFunc,
		timeFunc: timeFunc,
	}
}

type service struct {
	uuidFunc func() uuid.UUID
	timeFunc func() time.Time
}
