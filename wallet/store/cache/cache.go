package cache

import (
	"github.com/go-redis/redis"

	"github.com/sbuttigieg/test-quik-tech/wallet"
	"github.com/sbuttigieg/test-quik-tech/wallet/store"
)

// New create new redis store
func New(c *wallet.Config, db redis.UniversalClient) store.Cache {
	s := &cache{
		config: c,
		db:     db,
	}

	return s
}

type cache struct {
	config *wallet.Config
	db     redis.UniversalClient
}
