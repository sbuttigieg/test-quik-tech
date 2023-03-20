package api

import (
	"encoding/json"
	"fmt"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *service) Balance(walletID string) (float64, error) {
	var player models.Player

	u, ok := s.cache.GetKeyBytes(walletID)
	if !ok {
		fmt.Println("service balance", ok)
		// get from store
		// if not found => error "player not found"
		// if found store to cache
		// return store balance, nil
	}

	err := json.Unmarshal(u, &player)
	if err != nil {
		return 0, err
	}

	return player.Balance, nil
}
