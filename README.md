# test-quik-tech

## Steps to install application

`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/test-quik-tech ./cmd/main.go`

`docker network create test-quik-tech`

`docker build -t wallet-test-quik-tech .`

## To launch application

`docker-compose up -d`

## To stop application

`docker-compose down`

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
- Before balance, credit and debit endpoints are allowed, the user shall be verified by the auth endpoint. Once verified, the user will be marked as logged in in cache. If user is inactive for more than 1 minute, he will be logged out and auth will need to be repeated.
- The username and password are used as Basic Auth for all endpoints except auth.
- When a balance request is received, it is retrieved from cache. If it is not found in cache, then it will be retrieved from the store.
- When a debit request is received, if successful the balance of the player will be deducted by the debit amount. The balance is updated in both the store and cache.
- When a credit request is received, if successful the balance of the player will be increased by the debit amount. The balance is updated in both the store and cache.



## Endpoints

POST `/api/v1/wallets/:wallet_id/auth`  ==> to authenticate user and log him in

requires body in the following format:
```
{
    "username": "some user",
    "password': "some password"
}
```

returns 

POST `/api/v1/wallets/:wallet_id/` ==> to register a new user

requires body in the following format:
```
{
    "username": "some user",
    "password': "some password"
}
```

GET `/api/v1/wallets/balance` ==> to get user balance