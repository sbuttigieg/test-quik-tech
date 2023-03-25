# test-quik-tech

## Steps to install application

`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/test-quik-tech ./cmd/main.go`

`docker network create test-quik-tech`

`docker build -t wallet-test-quik-tech .`

## To launch application

`docker-compose up -d`

## To stop application

`docker-compose down`

## Unit tests

`Unit tests for the service layer can be found in branch feature/unit_tests. They have not been included in the master branch since they were finished after the deadline to complete the task.`

`These tests have a 100% coverage of the service layer and include mocking cache and mysql functions in the store layer.`

## Functionality
- Database contains 3 users with the following credentials
```
{
    "wallet_id": "6cc4ee0d-9919-4857-a70d-9b7283957e16",
    "username": "Bob",
    "password': "123456"
},
{
    "wallet_id": "0924f01f-3f70-4fe4-ac82-dce4b30e2a7f",
    "username": "Joe",
    "password': "654321"
},
{
    "wallet_id": "d2ba410a-9bc4-476b-86af-c55525b527df",
    "username": "Dave",
    "password': "456789"
}
```
- Before balance, credit and debit endpoints are allowed, the player shall be verified by the auth endpoint. If player is inactive for more than 5 minutes, no transactions will be allowed until auth request is repeated.
- The username and password are used as Basic Auth for all endpoints except auth.
- When a balance request is received, it is retrieved from cache. (during authentication the player data is retrieved from cache/store so there is certainty that the balance is in cache)
- When a debit request is received, if successful the balance of the player will be deducted by the debit amount. The balance is updated in both the store and cache.
- When a credit request is received, if successful the balance of the player will be increased by the credit amount. The balance is updated in both the store and cache.

## Endpoints

POST `/api/v1/wallets/:wallet_id/auth`  ==> to authenticate player and log him in

requires body in the following format:
```
{
    "username": "some username",
    "password': "some password"
}
```

returns response in the following format if successful:
```
{
    "wallet_id": "some walletID",
    "balance": a number
}
```

GET `/api/v1/wallets/:wallet_id/balance` ==> to get player balance

returns response in the following format if successful:
```
{
    "wallet_id": "some walletID",
    "balance": a number
}
```

POST `/api/v1/wallets/:wallet_id/credit` ==> to add funds to player balance

requires body in the following format:
```
{
    "amount": "some amount",
    "description': "some string"
}
```

returns response in the following format if successful:
```
{
    "id": "some transaction ID",
    "wallet_id": "some wallet ID",
    "amount": some amount,
    "type": "some string",
    "balance": some amount,
    "created_at": some date
}
```

POST `/api/v1/wallets/:wallet_id/debit` ==> to deduct funds from player balance

requires body in the following format:
```
{
    "amount": "some amount",
    "description': "some string"
}
```

returns response in the following format if successful:
```
{
    "id": "some transaction ID",
    "wallet_id": "some wallet ID",
    "amount": some amount,
    "type": "some string",
    "balance": some amount,
    "created_at": some date
}
```