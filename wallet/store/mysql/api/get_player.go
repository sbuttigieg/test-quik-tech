package api

import "github.com/sbuttigieg/test-quik-tech/wallet/models"

func (s *store) GetPlayer(walletID string) (*models.Player, error) {
	var player models.Player

	err := s.db.QueryRow("SELECT * FROM players WHERE wallet_id = ?;", walletID).Scan(
		&player.WalletID,
		&player.Balance,
		&player.Username,
		&player.Password,
		&player.LastActivity,
	)
	if err != nil {
		return nil, err
	}

	return &player, nil
}
