package cache

import (
	"errors"

	"github.com/go-redis/redis"
)

func (s *cache) GetKeyInt64(key string) (int64, bool) {
	value, err := s.db.Get(key).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, false
		}

		return 0, false
	}

	s.db.Expire(key, s.config.CacheExpiry)

	return value, true
}

func (s *cache) GetKeyString(key string) (string, bool) {
	value := s.db.Get(key).String()
	if value == "" {
		return "", false
	}

	s.db.Expire(key, s.config.CacheExpiry)

	return value, true
}

func (s *cache) GetKeyBytes(key string) ([]byte, bool) {
	value, err := s.db.Get(key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, false
		}

		return nil, false
	}

	s.db.Expire(key, s.config.CacheExpiry)

	return value, true
}
