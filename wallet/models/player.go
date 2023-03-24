package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Player struct {
	WalletID     string          `json:"wallet_id" gorm:"primaryKey"`
	Balance      decimal.Decimal `json:"balance"`
	Username     string          `json:"username"`
	Password     string          `json:"password"`
	LastActivity time.Time       `json:"last_activity"`
}
