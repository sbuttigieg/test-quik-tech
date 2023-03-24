package store

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func DBInit(db *gorm.DB) error {
	const startBalance = 100

	err := db.AutoMigrate(&models.Player{}, &models.Transaction{})
	if err != nil {
		return err
	}

	var count int64

	db.Table("players").Count(&count)

	if count == 0 {
		var players = []models.Player{
			{WalletID: "6cc4ee0d-9919-4857-a70d-9b7283957e16", Balance: decimal.NewFromInt(startBalance), Username: "Bob", Password: "123456"},
			{WalletID: "0924f01f-3f70-4fe4-ac82-dce4b30e2a7f", Balance: decimal.NewFromInt(startBalance), Username: "Joe", Password: "654321"},
			{WalletID: "d2ba410a-9bc4-476b-86af-c55525b527df", Balance: decimal.NewFromInt(startBalance), Username: "Dave", Password: "456789"},
		}

		db.Create(&players)
	}

	return nil
}
