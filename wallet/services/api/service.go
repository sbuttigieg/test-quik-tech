package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/sbuttigieg/test-quik-tech/wallet"
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
	"github.com/sbuttigieg/test-quik-tech/wallet/store"
)

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service
type Service interface {
	Auth(string, string, string, bool) (*models.Player, error)
	Balance(string) (*decimal.Decimal, error)
	Credit(string, string, decimal.Decimal) (*models.Transaction, error)
	Debit(string, string, decimal.Decimal) (*models.Transaction, error)
}

func New(config *wallet.Config, cache store.Cache, store Store, logger *logrus.Logger, uuidFunc func() uuid.UUID, timeFunc func() time.Time) Service {
	return &service{
		config:   config,
		cache:    cache,
		store:    store,
		logger:   logger,
		uuidFunc: uuidFunc,
		timeFunc: timeFunc,
	}
}

type service struct {
	config   *wallet.Config
	cache    store.Cache
	store    Store
	logger   *logrus.Logger
	uuidFunc func() uuid.UUID
	timeFunc func() time.Time
}
