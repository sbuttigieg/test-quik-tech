package api

import (
	"encoding/json"
	"errors"

	"github.com/shopspring/decimal"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *service) Balance(walletID string) (*decimal.Decimal, error) {
	var player *models.Player

	p, ok := s.cache.GetKeyBytes(walletID)
	if !ok {
		return nil, errors.New("player not found")
	}

	err := json.Unmarshal(p, &player)
	if err != nil {
		return nil, err
	}

	// Check if player is active
	elapsed := s.timeFunc().Sub(player.LastActivity)

	if elapsed >= s.config.SessionExpiry {
		return nil, errors.New("player not logged in")
	}

	// Update last activity
	player, err = s.store.ActivePlayer(walletID)
	if err != nil {
		return nil, err
	}

	// Can be done as store middleware
	err = s.cache.SetKey(walletID, player, s.config.CacheExpiry)
	if err != nil {
		return nil, err
	}

	return &player.Balance, nil
}
