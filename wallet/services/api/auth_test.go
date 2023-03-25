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
	cacheMocks "github.com/sbuttigieg/test-quik-tech/wallet/store/mocks"
)

func TestAuth(t *testing.T) {
	cases := []struct {
		it       string
		walletID string
		username string
		password string
		login    bool

		cacheGetInput          string
		cacheGetResponsePlayer []byte
		cacheGetResponseOK     bool

		storeGetPlayerInput          string
		storeGetPlayerResponsePlayer *models.Player
		storeGetPlayerResponseError  error

		cacheSetKeyInput      []string
		cacheSetValueInput    []interface{}
		cacheSetDurationInput []time.Duration
		cacheSetResponseError []error

		storeActivePlayerInput          string
		storeActivePlayerResponsePlayer *models.Player
		storeActivePlayerResponseError  error

		expectedError  string
		expectedResult *models.Player
	}{
		{
			it:       "it returns a player from cache for a non-login request",
			walletID: "mock-id",
			username: "mock-username",
			password: "mock-password",
			login:    false,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "username": "mock-username", "password": "mock-password"}`),
			cacheGetResponseOK:     true,

			expectedResult: &models.Player{
				WalletID: "mock-id",
				Username: "mock-username",
				Password: "mock-password",
			},
			expectedError: "",
		},
		{
			it:       "it returns a player from store for a login-request",
			walletID: "mock-id",
			username: "mock-username",
			password: "mock-password",
			login:    true,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(""),
			cacheGetResponseOK:     false,

			storeGetPlayerInput: "mock-id",
			storeGetPlayerResponsePlayer: &models.Player{
				WalletID: "mock-id",
				Username: "mock-username",
				Password: "mock-password",
			},
			storeGetPlayerResponseError: nil,

			cacheSetKeyInput: []string{"mock-id", "mock-id"},
			cacheSetValueInput: []interface{}{
				&models.Player{
					WalletID: "mock-id",
					Username: "mock-username",
					Password: "mock-password",
				},
				&models.Player{
					WalletID: "mock-id",
					Username: "mock-username",
					Password: "mock-password",
				},
			},
			cacheSetDurationInput: []time.Duration{10, 10},
			cacheSetResponseError: []error{nil, nil},

			storeActivePlayerInput: "mock-id",
			storeActivePlayerResponsePlayer: &models.Player{
				WalletID: "mock-id",
				Username: "mock-username",
				Password: "mock-password",
			},
			storeActivePlayerResponseError: nil,

			expectedResult: &models.Player{
				WalletID: "mock-id",
				Username: "mock-username",
				Password: "mock-password",
			},
			expectedError: "",
		},
		{
			it:       "it fails to store new activated player in cache",
			walletID: "mock-id",
			username: "mock-username",
			password: "mock-password",
			login:    true,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "username": "mock-username", "password": "mock-password"}`),
			cacheGetResponseOK:     true,

			cacheSetKeyInput: []string{"mock-id"},
			cacheSetValueInput: []interface{}{&models.Player{
				WalletID: "mock-id",
				Username: "mock-username",
				Password: "mock-password",
			}},
			cacheSetDurationInput: []time.Duration{10},
			cacheSetResponseError: []error{errors.New("mock-error")},

			storeActivePlayerInput: "mock-id",
			storeActivePlayerResponsePlayer: &models.Player{
				WalletID: "mock-id",
				Username: "mock-username",
				Password: "mock-password",
			},
			storeActivePlayerResponseError: nil,

			expectedResult: nil,
			expectedError:  "mock-error",
		},
		{
			it:       "it fails to activate player",
			walletID: "mock-id",
			username: "mock-username",
			password: "mock-password",
			login:    true,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id", "username": "mock-username", "password": "mock-password"}`),
			cacheGetResponseOK:     true,

			storeActivePlayerInput: "mock-id",
			storeActivePlayerResponsePlayer: &models.Player{
				WalletID: "mock-id",
				Username: "mock-username",
				Password: "mock-password",
			},
			storeActivePlayerResponseError: errors.New("mock-error"),

			expectedResult: nil,
			expectedError:  "mock-error",
		},
		{
			it:       "it fails to with incorrect credentials",
			walletID: "mock-id",
			username: "mock-username",
			password: "mock-password",
			login:    true,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(`{"wallet_id": "mock-id"}`),
			cacheGetResponseOK:     true,

			expectedResult: nil,
			expectedError:  "incorrect credentials",
		},
		{
			it:       "it fails to unmarshal player",
			walletID: "mock-id",
			username: "mock-username",
			password: "mock-password",
			login:    true,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(""),
			cacheGetResponseOK:     true,

			expectedResult: nil,
			expectedError:  "unexpected end of JSON input",
		},
		{
			it:       "it fails to save player in cache",
			walletID: "mock-id",
			username: "mock-username",
			password: "mock-password",
			login:    true,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(""),
			cacheGetResponseOK:     false,

			storeGetPlayerInput: "mock-id",
			storeGetPlayerResponsePlayer: &models.Player{
				WalletID: "mock-id",
				Username: "mock-username",
				Password: "mock-password",
			},
			storeGetPlayerResponseError: nil,

			cacheSetKeyInput: []string{"mock-id"},
			cacheSetValueInput: []interface{}{&models.Player{
				WalletID: "mock-id",
				Username: "mock-username",
				Password: "mock-password",
			}},
			cacheSetDurationInput: []time.Duration{10},
			cacheSetResponseError: []error{errors.New("mock-error")},

			expectedResult: nil,
			expectedError:  "mock-error",
		},
		{
			it:       "it fails to get player from store",
			walletID: "mock-id",
			username: "mock-username",
			password: "mock-password",
			login:    true,

			cacheGetInput:          "mock-id",
			cacheGetResponsePlayer: []byte(""),
			cacheGetResponseOK:     false,

			storeGetPlayerInput: "mock-id",
			storeGetPlayerResponsePlayer: &models.Player{
				WalletID: "mock-id",
				Username: "mock-username",
				Password: "mock-password",
			},
			storeGetPlayerResponseError: errors.New("mock-error"),

			expectedResult: nil,
			expectedError:  "player not found",
		},
		{
			it:       "it fails with missing credentials",
			walletID: "mock-id",
			username: "",
			password: "mock-password",
			login:    true,

			expectedResult: nil,
			expectedError:  "missing credentials",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			i := 0
			cache := &cacheMocks.CacheMock{
				GetKeyBytesFunc: func(s string) ([]byte, bool) {
					checkIs.Equal(s, tc.cacheGetInput)

					return tc.cacheGetResponsePlayer, tc.cacheGetResponseOK
				},
				SetKeyFunc: func(s string, val interface{}, duration time.Duration) error {
					defer func() {
						i++
					}()
					checkIs.Equal(s, tc.cacheSetKeyInput[i])
					checkIs.Equal(val, tc.cacheSetValueInput[i])
					checkIs.Equal(duration, tc.cacheSetDurationInput[i])

					return tc.cacheSetResponseError[i]
				},
			}

			store := &apiMocks.StoreMock{
				GetPlayerFunc: func(s string) (*models.Player, error) {
					checkIs.Equal(s, tc.storeGetPlayerInput)

					return tc.storeGetPlayerResponsePlayer, tc.storeGetPlayerResponseError
				},
				ActivePlayerFunc: func(s string) (*models.Player, error) {
					checkIs.Equal(s, tc.storeActivePlayerInput)

					return tc.storeActivePlayerResponsePlayer, tc.storeActivePlayerResponseError
				},
			}

			config := &wallet.Config{CacheExpiry: 10}
			service := api.New(config, cache, store, nil, nil, nil)

			player, err := service.Auth(tc.walletID, tc.username, tc.password, tc.login)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}

			checkIs.Equal(player, tc.expectedResult)
		})
	}
}
