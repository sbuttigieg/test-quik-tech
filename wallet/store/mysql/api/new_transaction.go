package api

import (
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (s *store) NewTransaction(transaction models.Transaction) (*models.Transaction, error) {
	var res models.Transaction

	_, err := s.db.Exec("INSERT INTO transactions (id, wallet_id, amount, type, balance) VALUES ( ?, ?, ?, ?, ? )",
		transaction.TransactionID,
		transaction.WalletID,
		transaction.Amount,
		transaction.Type,
		transaction.Balance,
	)
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow("SELECT * FROM transactions WHERE id = ?;", transaction.TransactionID).Scan(
		&res.TransactionID,
		&res.WalletID,
		&res.Amount,
		&res.Type,
		&res.Balance,
		&res.Created,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
