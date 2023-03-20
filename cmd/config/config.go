package config

import (
	"os"
	"strconv"
	"time"

	"github.com/sbuttigieg/test-quik-tech/wallet"
)

// NewConfig create new config
func NewConfig() (*wallet.Config, error) {
	env := os.Getenv("ENV")
	version := os.Getenv("VERSION")

	redisExpiry, err := strconv.Atoi(os.Getenv("REDIS_EXPIRY_SEC"))
	if err != nil {
		return nil, err
	}

	sessionExpiry, err := strconv.Atoi(os.Getenv("SESSION_EXPIRY_SEC"))
	if err != nil {
		return nil, err
	}

	c := &wallet.Config{
		Env:           env,
		Version:       version,
		CacheExpiry:   time.Duration(redisExpiry) * time.Second,
		SessionExpiry: time.Duration(sessionExpiry) * time.Second,
	}

	return c, nil
}
