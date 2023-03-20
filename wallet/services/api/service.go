package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/sbuttigieg/test-quik-tech/wallet"
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
	"github.com/sbuttigieg/test-quik-tech/wallet/store"
)

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service
type Service interface {
	Auth(string, string, string) (*models.User, error)
	Balance(string) (float64, error)
	Credit(string, float64) (float64, error)
	Debit(string, float64) (float64, error)
}

func New(config *wallet.Config, cache store.Cache, uuidFunc func() uuid.UUID, timeFunc func() time.Time) Service {
	return &service{
		config:   config,
		cache:    cache,
		uuidFunc: uuidFunc,
		timeFunc: timeFunc,
	}
}

type service struct {
	config   *wallet.Config
	cache    store.Cache
	uuidFunc func() uuid.UUID
	timeFunc func() time.Time
}
