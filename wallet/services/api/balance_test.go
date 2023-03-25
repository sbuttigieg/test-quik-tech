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

func TestBalance(t *testing.T) {
	cases := []struct {
		it       string
		walletID string

		cacheGetInput          string
		cacheGetResponsePlayer []byte
		cacheGetResponseOK     bool

		storeActivePlayerInput          string
		storeActivePlayerResponsePlayer *models.Player
		storeActivePlayerResponseError  error

		cacheSetKeyInput      string
		cacheSetValueInput    interface{}
		cacheSetDurationInput time.Duration
		cacheSetResponseError error

		expectedError  string
		expectedResult *decimal.Decimal
	}{
		{
			it:       "it returns a player balance",
			walletID: "mock-id",

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "balance": "0", "last_activity": "2020-01-02T03:04:05Z"}`),
			cacheGetResponseOK:     true,

			storeActivePlayerInput: "mock-id",
			storeActivePlayerResponsePlayer: &models.Player{
				WalletID:     "mock-id",
				Balance:      decimal.NewFromFloat(0),
				LastActivity: testsuite.GenTime(),
			},
			storeActivePlayerResponseError: nil,

			cacheSetKeyInput: "mock-id",
			cacheSetValueInput: &models.Player{
				WalletID:     "mock-id",
				Balance:      decimal.NewFromFloat(0),
				LastActivity: testsuite.GenTime(),
			},
			cacheSetDurationInput: 10,
			cacheSetResponseError: nil,
			expectedResult:        &decimal.Zero,
			expectedError:         "",
		},
		{
			it:       "it fails to store new activated player in cache",
			walletID: "mock-id",

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "balance": "0", "last_activity": "2020-01-02T03:04:05Z"}`),
			cacheGetResponseOK:     true,

			storeActivePlayerInput: "mock-id",
			storeActivePlayerResponsePlayer: &models.Player{
				WalletID:     "mock-id",
				Balance:      decimal.NewFromFloat(0),
				LastActivity: testsuite.GenTime(),
			},
			storeActivePlayerResponseError: nil,

			cacheSetKeyInput: "mock-id",
			cacheSetValueInput: &models.Player{
				WalletID:     "mock-id",
				Balance:      decimal.NewFromFloat(0),
				LastActivity: testsuite.GenTime(),
			},
			cacheSetDurationInput: 10,
			cacheSetResponseError: errors.New("mock-error"),

			expectedResult: nil,
			expectedError:  "mock-error",
		},
		{
			it:       "it fails to activate player",
			walletID: "mock-id",

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "balance": "0", "last_activity": "2020-01-02T03:04:05Z"}`),
			cacheGetResponseOK:     true,

			storeActivePlayerInput:          "mock-id",
			storeActivePlayerResponsePlayer: nil,
			storeActivePlayerResponseError:  errors.New("mock-error"),

			expectedResult: nil,
			expectedError:  "mock-error",
		},
		{
			it:       "it fails with player session expired",
			walletID: "mock-id",

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "balance": "0", "last_activity": "2020-01-02T03:01:05Z"}`),
			cacheGetResponseOK:     true,

			expectedResult: nil,
			expectedError:  "player not logged in",
		},
		{
			it:       "it fails to unmarshal player",
			walletID: "mock-id",

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(""),
			cacheGetResponseOK:     true,

			expectedResult: nil,
			expectedError:  "unexpected end of JSON input",
		},
		{
			it:       "it fails to get player from cache",
			walletID: "mock-id",

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(""),
			cacheGetResponseOK:     false,

			expectedResult: nil,
			expectedError:  "player not found",
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
				ActivePlayerFunc: func(s string) (*models.Player, error) {
					checkIs.Equal(s, tc.storeActivePlayerInput)

					return tc.storeActivePlayerResponsePlayer, tc.storeActivePlayerResponseError
				},
			}

			config := &wallet.Config{SessionExpiry: 100000000000, CacheExpiry: 10}
			service := api.New(config, cache, store, nil, nil, testsuite.GenTime)

			balance, err := service.Balance(tc.walletID)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			} else {
				checkIs.True(balance.Equal(*tc.expectedResult))
			}
		})
	}
}
