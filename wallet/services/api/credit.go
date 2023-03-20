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

	var player models.Player

	u, ok := s.cache.GetKeyBytes(walletID)
	if !ok {
		fmt.Println("service credit", ok)
		// get from store
		// if not found => error "player not found"
		// if found store to cache
		// set player to store player
	}

	if ok {
		err := json.Unmarshal(u, &player)
		if err != nil {
			return 0, err
		}
	}

	player.Balance += amount

	// Replace with store update (balance, last activity) that also includes cache update
	err := s.cache.SetKey(walletID, player, s.config.CacheExpiry)
	if err != nil {
		return 0, err
	}
	// *********************

	return player.Balance, nil
}
