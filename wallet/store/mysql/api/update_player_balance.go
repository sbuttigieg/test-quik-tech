package api

import (
	"github.com/shopspring/decimal"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *store) UpdatePlayerBalance(walletID string, amount decimal.Decimal) (*decimal.Decimal, error) {
	result := s.db.Model(&models.Player{}).Where("wallet_id = ?", walletID).Update("balance", amount)
	if result.Error != nil {
		return nil, result.Error
	}

	player, err := s.GetPlayer(walletID)
	if err != nil {
		return nil, err
	}

	return &player.Balance, nil
}
