package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/sbuttigieg/test-quik-tech/wallet"
	service "github.com/sbuttigieg/test-quik-tech/wallet/services/api"
	"github.com/sbuttigieg/test-quik-tech/wallet/store"
)

func NewService(cfg *wallet.Config, cache store.Cache, uuidFunc func() uuid.UUID, timeFunc func() time.Time) (service.Service, error) {
	service := service.New(cfg, cache, uuidFunc, timeFunc)

	return service, nil
}
