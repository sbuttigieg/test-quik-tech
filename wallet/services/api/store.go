package api

import (
	"github.com/shopspring/decimal"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

//go:generate moq -out ./mocks/store.go -pkg mocks  . Store
type Store interface {
	// players
	ActivePlayer(string) (*models.Player, error)
	GetPlayer(string) (*models.Player, error)
	UpdatePlayerBalance(string, decimal.Decimal) (*decimal.Decimal, error)

	// transactions
	NewTransaction(models.Transaction) (*models.Transaction, error)
}
