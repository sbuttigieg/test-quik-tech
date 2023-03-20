package models

import "github.com/shopspring/decimal"

type Balance struct {
	WalletID string          `json:"wallet_id"`
	Balance  decimal.Decimal `json:"balance"`
}
