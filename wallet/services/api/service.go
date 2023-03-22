package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/sbuttigieg/test-quik-tech/wallet"
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
	"github.com/sbuttigieg/test-quik-tech/wallet/store"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service
type Service interface {
	Auth(string, string, string) (*models.Player, error)
	Balance(string) (*decimal.Decimal, error)
	Credit(string, decimal.Decimal) (*decimal.Decimal, error)
	Debit(string, decimal.Decimal) (*decimal.Decimal, error)
}

func New(config *wallet.Config, cache store.Cache, logger *logrus.Logger, uuidFunc func() uuid.UUID, timeFunc func() time.Time) Service {
	return &service{
		config:   config,
		cache:    cache,
		logger:   logger,
		uuidFunc: uuidFunc,
		timeFunc: timeFunc,
	}
}

type service struct {
	config   *wallet.Config
	cache    store.Cache
	logger   *logrus.Logger
	uuidFunc func() uuid.UUID
	timeFunc func() time.Time
}
