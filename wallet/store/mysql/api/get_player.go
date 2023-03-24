package api

import "github.com/sbuttigieg/test-quik-tech/wallet/models"

func (s *store) GetPlayer(walletID string) (*models.Player, error) {
	var player models.Player

	result := s.db.First(&player, "wallet_id = ?", walletID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &player, nil
}
