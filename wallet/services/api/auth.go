package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *service) Auth(walletID, username, password string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("missing credentials")
	}

	var user models.User

	u, ok := s.cache.GetKeyBytes(walletID)
	if !ok {
		fmt.Println("service Auth", ok)
		// get from store
		// if not found => error "user not found"
		// if found store to cache
		// set user to store user
	}

	if ok {
		err := json.Unmarshal(u, &user)
		if err != nil {
			return nil, err
		}
	}

	if user.Username != username || user.Password != password {
		return nil, errors.New("incorrect credentials")
	}

	// update last activity

	return &user, nil
}
