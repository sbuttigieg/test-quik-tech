package api

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *service) Debit(walletID string, description string, amount decimal.Decimal) (*models.Transaction, error) {
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
	elapsed := time.Since(player.LastActivity)

	if elapsed >= s.config.SessionExpiry {
		return nil, errors.New("player not logged in")
	}

	res := player.Balance.Sub(amount)
	if res.LessThan(decimal.Zero) {
		return nil, errors.New("insufficient funds")
	}

	balance, err := s.store.UpdatePlayerBalance(walletID, res)
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
		ID:       uuid.New().String(),
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
