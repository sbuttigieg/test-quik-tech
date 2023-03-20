package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *service) Auth(walletID, username, password string) (*models.Player, error) {
	if username == "" || password == "" {
		return nil, errors.New("missing credentials")
	}

	var player models.Player

	u, ok := s.cache.GetKeyBytes(walletID)
	if !ok {
		fmt.Println("service Auth", ok)
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

	if player.Username != username || player.Password != password {
		return nil, errors.New("incorrect credentials")
	}

	// update last activity

	return &player, nil
}
