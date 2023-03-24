package api

import (
	"context"
	"time"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *store) ActivePlayer(walletID string) (*models.Player, error) {
	_, err := s.db.ExecContext(context.Background(), "UPDATE players SET last_activity = ? WHERE wallet_id = ?", time.Now(), walletID)
	if err != nil {
		return nil, err
	}

	player, err := s.GetPlayer(walletID)
	if err != nil {
		return nil, err
	}

	return player, nil
}
