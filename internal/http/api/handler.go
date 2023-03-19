package api

import "github.com/sbuttigieg/test-quik-tech/internal/services/api"

func New(serv api.Service) *Handler {
	return &Handler{
		service: serv,
	}
}

type Handler struct {
	service api.Service
}
