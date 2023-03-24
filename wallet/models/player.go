package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Player struct {
	WalletID     string          `json:"wallet_id" gorm:"type:varchar(50);primaryKey;not null;unique"`
	Balance      decimal.Decimal `json:"balance" gorm:"type:decimal(15,10);"`
	Username     string          `json:"username" gorm:"type:text;"`
	Password     string          `json:"password" gorm:"type:text;"`
	LastActivity time.Time       `json:"last_activity" gorm:"type:TIMESTAMP;default:current_timestamp;autoUpdateTime"`
}
