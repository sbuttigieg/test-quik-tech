package api

import (
	"github.com/sbuttigieg/test-quik-tech/wallet/services/api"
	store "github.com/sbuttigieg/test-quik-tech/wallet/store/mysql/api"
	"gorm.io/gorm"
)

func NewStore(db *gorm.DB) api.Store {
	return store.New(db)
}
