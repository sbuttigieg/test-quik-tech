package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	TransactionID string          `json:"id" gorm:"primaryKey"`
	WalletID      string          `json:"wallet_id"`
	Amount        decimal.Decimal `json:"amount"`
	Type          string          `json:"type"`
	Balance       decimal.Decimal `json:"balance"`
	Created       time.Time       `json:"created_at"`
}
