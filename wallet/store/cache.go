package store

import (
	"time"
)

// Cache interface
type Cache interface {
	SetKey(string, interface{}, time.Duration) error
	GetKeyInt64(string) (int64, bool)
	GetKeyString(string) (string, bool)
	GetKeyBytes(string) ([]byte, bool)
}
