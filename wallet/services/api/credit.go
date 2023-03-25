package api

import (
	"encoding/json"
	"errors"

	"github.com/shopspring/decimal"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *service) Credit(walletID string, description string, amount decimal.Decimal) (*models.Transaction, error) {
	if amount.LessThan(decimal.Zero) {
		return nil, errors.New("negative value")
	}

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

	balance, err := s.store.UpdatePlayerBalance(walletID, player.Balance.Add(amount))
	if err != nil {
		return nil, err
	}

	player.Balance = *balance

	// Can be done as store middleware
	err = s.cache.SetKey(walletID, player, s.config.CacheExpiry)
	if err != nil {
		return nil, err
	}

	transaction, err := s.store.NewTransaction(models.Transaction{
		ID:       s.uuidFunc().String(),
		WalletID: walletID,
		Amount:   amount,
		Type:     description,
		Balance:  player.Balance,
	})
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
