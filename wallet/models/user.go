package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Player struct {
	WalletID     string
	Balance      decimal.Decimal
	Username     string
	Password     string
	LastActivity time.Time
}
