package api

import "github.com/sbuttigieg/test-quik-tech/wallet/services/api"

const (
	IncorrectCredentials    = "incorrect credentials"
	MissingCredentialsError = "missing credentials"
	NegativeValueError      = "negative value"
	InsufficientFundsError  = "insufficient funds"
	PlayerNotFoundError     = "player not found"
	PlayerNotLoggedIn       = "player not logged in"
)

func New(serv api.Service) *Handler {
	return &Handler{
		service: serv,
	}
}

type Handler struct {
	service api.Service
}
