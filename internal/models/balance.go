package models

type Balance struct {
	WalletID string  `json:"wallet_id"`
	Balance  float64 `json:"balance"`
}
