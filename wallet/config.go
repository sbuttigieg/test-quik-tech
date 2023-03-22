package wallet

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Env           string
	Version       string
	CacheExpiry   time.Duration
	SessionExpiry time.Duration
	LogLevel      logrus.Level
}
