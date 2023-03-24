package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID        string          `json:"id" gorm:"type:varchar(50);primaryKey;not null;unique"`
	WalletID  string          `json:"wallet_id" gorm:"type:text;"`
	Amount    decimal.Decimal `json:"amount" gorm:"type:decimal(15,10);"`
	Type      string          `json:"type" gorm:"type:text;"`
	Balance   decimal.Decimal `json:"balance" gorm:"type:decimal(15,10);"`
	CreatedAt time.Time       `json:"created_at" gorm:"type:TIMESTAMP;autoUpdateTime"`
}
