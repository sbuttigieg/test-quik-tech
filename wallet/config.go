package wallet

import (
	"time"
)

type Config struct {
	Env           string
	Version       string
	CacheExpiry   time.Duration
	SessionExpiry time.Duration
}
