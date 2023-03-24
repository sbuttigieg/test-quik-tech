package api

import (
	"time"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *store) ActivePlayer(walletID string) (*models.Player, error) {
	result := s.db.Model(&models.Player{}).Where("wallet_id = ?", walletID).Update("last_activity", time.Now())
	if result.Error != nil {
		return nil, result.Error
	}

	player, err := s.GetPlayer(walletID)
	if err != nil {
		return nil, err
	}

	return player, nil
}
