package api

import (
	"encoding/json"
	"fmt"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *service) Balance(walletID string) (float64, error) {
	var user models.User

	u, ok := s.cache.GetKeyBytes(walletID)
	if !ok {
		fmt.Println("service balance", ok)
		// get from store
		// if not found => error "user not found"
		// if found store to cache
		// return store balance, nil
	}

	err := json.Unmarshal(u, &user)
	if err != nil {
		return 0, err
	}

	return user.Balance, nil
}
