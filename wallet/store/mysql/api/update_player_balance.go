package api

import (
	"context"

	"github.com/shopspring/decimal"
)

func (s *store) UpdatePlayerBalance(walletID string, amount decimal.Decimal) (*decimal.Decimal, error) {
	_, err := s.db.ExecContext(context.Background(), "UPDATE players SET balance = ? WHERE wallet_id = ?", amount, walletID)
	if err != nil {
		return nil, err
	}

	player, err := s.GetPlayer(walletID)
	if err != nil {
		return nil, err
	}

	return &player.Balance, nil
}
