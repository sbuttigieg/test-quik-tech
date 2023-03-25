package api_test

import (
	"errors"
	"testing"
	"time"

	"github.com/matryer/is"
	"github.com/sbuttigieg/test-quik-tech/wallet"
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
	"github.com/sbuttigieg/test-quik-tech/wallet/services/api"
	apiMocks "github.com/sbuttigieg/test-quik-tech/wallet/services/api/mocks"
	"github.com/sbuttigieg/test-quik-tech/wallet/services/testsuite"
	cacheMocks "github.com/sbuttigieg/test-quik-tech/wallet/store/mocks"
	"github.com/shopspring/decimal"
)

func TestDebit(t *testing.T) {
	cases := []struct {
		it          string
		walletID    string
		description string
		amount      decimal.Decimal

		cacheGetInput          string
		cacheGetResponsePlayer []byte
		cacheGetResponseOK     bool

		storeUpdateBalanceInputID       string
		storeUpdateBalanceInputAmount   decimal.Decimal
		storeUpdateBalanceResponse      *decimal.Decimal
		storeUpdateBalanceResponseError error

		cacheSetKeyInput      string
		cacheSetValueInput    interface{}
		cacheSetDurationInput time.Duration
		cacheSetResponseError error

		storeNewTransactionInput         models.Transaction
		storeNewTransactionResponse      *models.Transaction
		storeNewTransactionResponseError error

		expectedError  string
		expectedResult *models.Transaction
	}{
		{
			it:          "it returns a transaction",
			walletID:    "mock-id",
			description: "mock-desc",
			amount:      decimal.Zero,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "balance": "0", "last_activity": "2020-01-02T03:04:05Z"}`),
			cacheGetResponseOK:     true,

			storeUpdateBalanceInputID:       "mock-id",
			storeUpdateBalanceInputAmount:   decimal.Zero,
			storeUpdateBalanceResponse:      &decimal.Zero,
			storeUpdateBalanceResponseError: nil,

			cacheSetKeyInput: "mock-id",
			cacheSetValueInput: &models.Player{
				WalletID:     "mock-id",
				Balance:      decimal.Zero,
				LastActivity: testsuite.GenTime(),
			},
			cacheSetDurationInput: 10,
			cacheSetResponseError: nil,

			storeNewTransactionInput: models.Transaction{
				ID:       testsuite.GenUUID().String(),
				WalletID: "mock-id",
				Amount:   decimal.Zero,
				Type:     "mock-desc",
				Balance:  decimal.Zero,
			},
			storeNewTransactionResponse: &models.Transaction{
				ID:       testsuite.GenUUID().String(),
				WalletID: "mock-id",
				Amount:   decimal.Zero,
				Type:     "mock-desc",
				Balance:  decimal.Zero,
			},
			storeNewTransactionResponseError: nil,

			expectedResult: &models.Transaction{
				ID:       testsuite.GenUUID().String(),
				WalletID: "mock-id",
				Amount:   decimal.Zero,
				Type:     "mock-desc",
				Balance:  decimal.Zero,
			},
			expectedError: "",
		},
		{
			it:          "it fails to save a new transaction in store",
			walletID:    "mock-id",
			description: "mock-desc",
			amount:      decimal.Zero,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "balance": "0", "last_activity": "2020-01-02T03:04:05Z"}`),
			cacheGetResponseOK:     true,

			storeUpdateBalanceInputID:       "mock-id",
			storeUpdateBalanceInputAmount:   decimal.Zero,
			storeUpdateBalanceResponse:      &decimal.Zero,
			storeUpdateBalanceResponseError: nil,

			cacheSetKeyInput: "mock-id",
			cacheSetValueInput: &models.Player{
				WalletID:     "mock-id",
				Balance:      decimal.Zero,
				LastActivity: testsuite.GenTime(),
			},
			cacheSetDurationInput: 10,
			cacheSetResponseError: nil,

			storeNewTransactionInput: models.Transaction{
				ID:       testsuite.GenUUID().String(),
				WalletID: "mock-id",
				Amount:   decimal.Zero,
				Type:     "mock-desc",
				Balance:  decimal.Zero,
			},
			storeNewTransactionResponse:      nil,
			storeNewTransactionResponseError: errors.New("mock-error"),

			expectedResult: nil,
			expectedError:  "mock-error",
		},
		{
			it:          "it fails to store updated player in cache",
			walletID:    "mock-id",
			description: "mock-desc",
			amount:      decimal.Zero,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "balance": "0", "last_activity": "2020-01-02T03:04:05Z"}`),
			cacheGetResponseOK:     true,

			storeUpdateBalanceInputID:       "mock-id",
			storeUpdateBalanceInputAmount:   decimal.Zero,
			storeUpdateBalanceResponse:      &decimal.Zero,
			storeUpdateBalanceResponseError: nil,

			cacheSetKeyInput: "mock-id",
			cacheSetValueInput: &models.Player{
				WalletID:     "mock-id",
				Balance:      decimal.Zero,
				LastActivity: testsuite.GenTime(),
			},
			cacheSetDurationInput: 10,
			cacheSetResponseError: errors.New("mock-error"),

			expectedResult: nil,
			expectedError:  "mock-error",
		},
		{
			it:          "it fails to update player balance in store",
			walletID:    "mock-id",
			description: "mock-desc",
			amount:      decimal.Zero,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "balance": "0", "last_activity": "2020-01-02T03:04:05Z"}`),
			cacheGetResponseOK:     true,

			storeUpdateBalanceInputID:       "mock-id",
			storeUpdateBalanceInputAmount:   decimal.Zero,
			storeUpdateBalanceResponse:      nil,
			storeUpdateBalanceResponseError: errors.New("mock-error"),

			expectedResult: nil,
			expectedError:  "mock-error",
		},
		{
			it:          "it fails with insufficient funds",
			walletID:    "mock-id",
			description: "mock-desc",
			amount:      decimal.Zero,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "balance": "-1", "last_activity": "2020-01-02T03:04:05Z"}`),
			cacheGetResponseOK:     true,

			storeUpdateBalanceInputID:       "mock-id",
			storeUpdateBalanceInputAmount:   decimal.Zero,
			storeUpdateBalanceResponse:      nil,
			storeUpdateBalanceResponseError: errors.New("mock-error"),

			expectedResult: nil,
			expectedError:  "insufficient funds",
		},
		{
			it:          "it fails with player session expired",
			walletID:    "mock-id",
			description: "mock-desc",
			amount:      decimal.Zero,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "balance": "0", "last_activity": "2020-01-02T03:01:05Z"}`),
			cacheGetResponseOK:     true,

			expectedResult: nil,
			expectedError:  "player not logged in",
		},
		{
			it:          "it fails to unmarshal player",
			walletID:    "mock-id",
			description: "mock-desc",
			amount:      decimal.Zero,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(""),
			cacheGetResponseOK:     true,

			expectedResult: nil,
			expectedError:  "unexpected end of JSON input",
		},
		{
			it:          "it fails to get player from cache",
			walletID:    "mock-id",
			description: "mock-desc",
			amount:      decimal.Zero,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(""),
			cacheGetResponseOK:     false,

			expectedResult: nil,
			expectedError:  "player not found",
		},
		{
			it:          "it fails to get player from cache",
			walletID:    "mock-id",
			description: "mock-desc",
			amount:      decimal.NewFromInt(-1),

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(""),
			cacheGetResponseOK:     false,

			expectedResult: nil,
			expectedError:  "negative value",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			cache := &cacheMocks.CacheMock{
				GetKeyBytesFunc: func(s string) ([]byte, bool) {
					checkIs.Equal(s, tc.cacheGetInput)

					return tc.cacheGetResponsePlayer, tc.cacheGetResponseOK
				},
				SetKeyFunc: func(s string, val interface{}, duration time.Duration) error {
					checkIs.Equal(s, tc.cacheSetKeyInput)
					checkIs.Equal(val, tc.cacheSetValueInput)
					checkIs.Equal(duration, tc.cacheSetDurationInput)

					return tc.cacheSetResponseError
				},
			}

			store := &apiMocks.StoreMock{
				UpdatePlayerBalanceFunc: func(s string, amount decimal.Decimal) (*decimal.Decimal, error) {
					checkIs.Equal(s, tc.storeUpdateBalanceInputID)
					checkIs.True(amount.Equal(tc.storeUpdateBalanceInputAmount))

					return tc.storeUpdateBalanceResponse, tc.storeUpdateBalanceResponseError
				},
				NewTransactionFunc: func(transaction models.Transaction) (*models.Transaction, error) {
					checkIs.Equal(transaction, tc.storeNewTransactionInput)

					return tc.storeNewTransactionResponse, tc.storeNewTransactionResponseError
				},
			}

			config := &wallet.Config{SessionExpiry: 100000000000, CacheExpiry: 10}
			service := api.New(config, cache, store, nil, testsuite.GenUUID, testsuite.GenTime)

			transaction, err := service.Debit(tc.walletID, tc.description, tc.amount)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(transaction, tc.expectedResult)

		})
	}
}
