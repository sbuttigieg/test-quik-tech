package middleware

import (
	"fmt"
	"time"

	"github.com/sbuttigieg/test-quik-tech/wallet/models"
	"github.com/sbuttigieg/test-quik-tech/wallet/services/api"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type loggingMiddleware struct {
	next    api.Service
	service string
	log     *logrus.Logger
}

// NewLoggingMiddleware creates a new logging middleware.
func NewLoggingMiddleware(next api.Service, logger *logrus.Logger) api.Service {
	m := loggingMiddleware{
		next:    next,
		service: "api",
		log:     logger,
	}

	return &m
}

func (m *loggingMiddleware) Auth(walletID, username, password string) (*models.Player, error) {
	start := time.Now()
	player, err := m.next.Auth(walletID, username, password)
	end := time.Now()

	logMsg := models.LogService{
		Layer:       "service",
		Duration:    end.Sub(start).String(),
		ServiceName: m.service,
		Method:      "Auth",
	}

	switch err {
	case nil:
		logMsg.Data = fmt.Sprintf("wallet_id: %v", player.WalletID)
	default:
		logMsg.Data = fmt.Sprintf("error: %v", err.Error())

	}

	m.log.Debug(logMsg)

	return player, err
}

func (m *loggingMiddleware) Balance(walletID string) (*decimal.Decimal, error) {
	start := time.Now()
	balance, err := m.next.Balance(walletID)
	end := time.Now()

	logMsg := models.LogService{
		Layer:       "service",
		Duration:    end.Sub(start).String(),
		ServiceName: m.service,
		Method:      "Balance",
	}

	switch err {
	case nil:
		logMsg.Data = fmt.Sprintf("wallet_id: %v, balance: %v", walletID, balance)
	default:
		logMsg.Data = fmt.Sprintf("error: %v", err.Error())
	}

	m.log.Debug(logMsg)

	return balance, err
}

func (m *loggingMiddleware) Credit(walletID string, amount decimal.Decimal) (*decimal.Decimal, error) {
	start := time.Now()
	balance, err := m.next.Credit(walletID, amount)
	end := time.Now()

	logMsg := models.LogService{
		Layer:       "service",
		Duration:    end.Sub(start).String(),
		ServiceName: m.service,
		Method:      "Credit",
	}

	switch err {
	case nil:
		logMsg.Data = fmt.Sprintf("wallet_id: %v, amount:, %v, balance: %v", walletID, amount, balance)
	default:
		logMsg.Data = fmt.Sprintf("error: %v", err.Error())
	}

	m.log.Debug(logMsg)

	return balance, err
}

func (m *loggingMiddleware) Debit(walletID string, amount decimal.Decimal) (*decimal.Decimal, error) {
	start := time.Now()
	balance, err := m.next.Debit(walletID, amount)
	end := time.Now()

	logMsg := models.LogService{
		Layer:       "service",
		Duration:    end.Sub(start).String(),
		ServiceName: m.service,
		Method:      "Debit",
	}

	switch err {
	case nil:
		logMsg.Data = fmt.Sprintf("wallet_id: %v, amount:, %v, balance: %v", walletID, amount, balance)
	default:
		logMsg.Data = fmt.Sprintf("error: %v", err.Error())
	}

	m.log.Debug(logMsg)

	return balance, err
}
