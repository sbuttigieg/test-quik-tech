package models

type Transaction struct {
	WalletID string  `json:"wallet_id"`
	Amount   float64 `json:"amount"`
	Type     string  `json:"transaction_type"`
	Balance  float64 `json:"balance"`
}
