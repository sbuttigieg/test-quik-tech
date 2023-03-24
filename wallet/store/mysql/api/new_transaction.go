package api

import (
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *store) NewTransaction(transaction models.Transaction) (*models.Transaction, error) {
	var res models.Transaction

	result := s.db.Create(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}

	result = s.db.First(&res, "id = ?", transaction.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &res, nil
}
