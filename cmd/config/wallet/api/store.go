package api

import (
	"database/sql"

	"github.com/sbuttigieg/test-quik-tech/wallet/services/api"
	store "github.com/sbuttigieg/test-quik-tech/wallet/store/mysql/api"
)

func NewStore(db *sql.DB) api.Store {
	return store.New(db)
}
