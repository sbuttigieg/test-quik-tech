package api

import (
	"github.com/sbuttigieg/test-quik-tech/internal/http/api"
	service "github.com/sbuttigieg/test-quik-tech/internal/services/api"
)

func NewHandlers(service service.Service) (*api.Handler, error) {
	return api.New(service), nil
}
