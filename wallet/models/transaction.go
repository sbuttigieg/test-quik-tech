package models

import "github.com/shopspring/decimal"

type Transaction struct {
	WalletID string          `json:"wallet_id"`
	Amount   decimal.Decimal `json:"amount"`
	Type     string          `json:"transaction_type"`
	Balance  decimal.Decimal `json:"balance"`
}
