package api

import (
	"github.com/sbuttigieg/test-quik-tech/wallet/http/api"
	service "github.com/sbuttigieg/test-quik-tech/wallet/services/api"
)

func NewHandlers(service service.Service) *api.Handler {
	return api.New(service)
}
