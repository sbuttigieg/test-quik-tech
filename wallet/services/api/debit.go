package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
	"github.com/shopspring/decimal"
)

func (s *service) Debit(walletID string, amount decimal.Decimal) (*decimal.Decimal, error) {
	if amount.LessThan(decimal.Zero) {
		return nil, errors.New("negative value")
	}

	var player models.Player

	u, ok := s.cache.GetKeyBytes(walletID)
	if !ok {
		fmt.Println("service debit", ok)
		// get from store
		// if not found => error "player not found"
		// if found store to cache
		// set player to store player
	}

	if ok {
		err := json.Unmarshal(u, &player)
		if err != nil {
			return nil, err
		}
	}

	res := player.Balance.Sub(amount)
	if res.LessThan(decimal.Zero) {
		return nil, errors.New("insufficient funds")
	}

	player.Balance = res

	// Replace with store update (balance, last activity) that also includes cache update
	err := s.cache.SetKey(walletID, player, s.config.CacheExpiry)
	if err != nil {
		return nil, err
	}
	// *********************

	return &player.Balance, nil
}
