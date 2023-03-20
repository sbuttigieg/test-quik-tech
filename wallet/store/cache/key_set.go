package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// GetMap return map by key
func (s *cache) SetKey(key string, value interface{}, expire time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}

	err = s.db.Set(key, data, expire).Err()
	if err != nil {
		return errors.Wrapf(err, "set cache")
	}

	return nil
}
