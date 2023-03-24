package api

import (
	"encoding/json"
	"errors"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *service) Auth(walletID, username, password string, login bool) (*models.Player, error) {
	if username == "" || password == "" {
		return nil, errors.New("missing credentials")
	}

	var player *models.Player

	var err error

	p, ok := s.cache.GetKeyBytes(walletID)
	if !ok {
		player, err = s.store.GetPlayer(walletID)
		if err != nil {
			return nil, errors.New("player not found")
		}

		// Can be done as store middleware
		err = s.cache.SetKey(walletID, player, s.config.CacheExpiry)
		if err != nil {
			return nil, err
		}
	}

	if ok {
		err := json.Unmarshal(p, &player)
		if err != nil {
			return nil, err
		}
	}

	if player.Username != username || player.Password != password {
		return nil, errors.New("incorrect credentials")
	}

	// Update last activity when auth func is called for login
	if login {
		player, err = s.store.ActivePlayer(walletID)
		if err != nil {
			return nil, err
		}

		// Can be done as store middleware
		err = s.cache.SetKey(walletID, player, s.config.CacheExpiry)
		if err != nil {
			return nil, err
		}
	}

	return player, nil
}
