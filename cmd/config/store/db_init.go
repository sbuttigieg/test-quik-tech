package store

import (
	"database/sql"
)

func DBInit(store *sql.DB) error {
	_, err := store.Exec("CREATE TABLE IF NOT EXISTS players (wallet_id VARCHAR(50) NOT NULL UNIQUE, balance DECIMAL(15,10), username TEXT, password TEXT, last_activity TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP);")
	if err != nil {
		return err
	}

	_, err = store.Exec("CREATE TABLE IF NOT EXISTS transactions (id VARCHAR(50) NOT NULL UNIQUE, wallet_id TEXT REFERENCES players(wallet_id), amount DECIMAL(15,10), type TEXT, balance DECIMAL(15,10), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP);")
	if err != nil {
		return err
	}

	var count int

	err = store.QueryRow("SELECT EXISTS (SELECT 1 FROM players);").Scan((&count))
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = store.Exec("INSERT INTO players (wallet_id, balance, username, password) VALUES('6cc4ee0d-9919-4857-a70d-9b7283957e16',100,'Bob','123456');")
		if err != nil {
			return err
		}

		_, err = store.Exec("INSERT INTO players (wallet_id, balance, username, password) VALUES('0924f01f-3f70-4fe4-ac82-dce4b30e2a7f',100,'Joe','654321');")
		if err != nil {
			return err
		}

		_, err = store.Exec("INSERT INTO players (wallet_id, balance, username, password) VALUES('d2ba410a-9bc4-476b-86af-c55525b527df',100,'Dave','456789');")
		if err != nil {
			return err
		}
	}

	return nil
}
