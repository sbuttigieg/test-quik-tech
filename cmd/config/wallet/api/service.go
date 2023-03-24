package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/sbuttigieg/test-quik-tech/wallet"
	service "github.com/sbuttigieg/test-quik-tech/wallet/services/api"
	"github.com/sbuttigieg/test-quik-tech/wallet/services/api/middleware"
	"github.com/sbuttigieg/test-quik-tech/wallet/store"
	"github.com/sirupsen/logrus"
)

func NewService(cfg *wallet.Config, cache store.Cache, store service.Store, logger *logrus.Logger, uuidFunc func() uuid.UUID, timeFunc func() time.Time) service.Service {
	service := service.New(cfg, cache, store, logger, uuidFunc, timeFunc)
	service = middleware.NewLoggingMiddleware(service, logger)

	return service
}
