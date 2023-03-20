package store

import (
	"github.com/go-redis/redis"

	"github.com/sbuttigieg/test-quik-tech/wallet"
	"github.com/sbuttigieg/test-quik-tech/wallet/store"
	"github.com/sbuttigieg/test-quik-tech/wallet/store/cache"
)

// NewInmem create new in memory store
func NewCache(cfg *wallet.Config, redis redis.UniversalClient) (store.Cache, error) {
	return cache.New(cfg, redis), nil
}
