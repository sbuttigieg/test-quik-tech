package models

import "time"

type User struct {
	WalletID     string
	Balance      float64
	Username     string
	Password     string
	LastActivity time.Time
}
