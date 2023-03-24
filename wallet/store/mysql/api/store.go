package api

import (
	"database/sql"

	"github.com/sbuttigieg/test-quik-tech/wallet/services/api"
)

func New(db *sql.DB) api.Store {
	s := &store{
		db:     db,
		models: "api",
	}

	return s
}

type store struct {
	db     *sql.DB
	models string
}
