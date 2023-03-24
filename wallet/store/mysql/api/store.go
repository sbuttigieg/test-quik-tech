package api

import (
	"gorm.io/gorm"

	"github.com/sbuttigieg/test-quik-tech/wallet/services/api"
)

func New(db *gorm.DB) api.Store {
	s := &store{
		db:     db,
		models: "api",
	}

	return s
}

type store struct {
	db     *gorm.DB
	models string
}
