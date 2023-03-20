package models

import "time"

type Player struct {
	WalletID     string
	Balance      float64
	Username     string
	Password     string
	LastActivity time.Time
}
