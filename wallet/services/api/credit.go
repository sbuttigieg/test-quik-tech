package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *service) Credit(walletID string, amount float64) (float64, error) {
	if amount < 0 {
		return 0, errors.New("negative value")
	}

	var user models.User

	u, ok := s.cache.GetKeyBytes(walletID)
	if !ok {
		fmt.Println("service credit", ok)
		// get from store
		// if not found => error "user not found"
		// if found store to cache
		// set user to store user
	}

	if ok {
		err := json.Unmarshal(u, &user)
		if err != nil {
			return 0, err
		}
	}

	user.Balance += amount

	// Replace with store update (balance, last activity) that also includes cache update
	err := s.cache.SetKey(walletID, user, s.config.CacheExpiry)
	if err != nil {
		return 0, err
	}
	// *********************

	return user.Balance, nil
}