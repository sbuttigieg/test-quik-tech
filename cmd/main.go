package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/sbuttigieg/test-quik-tech/cmd/config"
	"github.com/sbuttigieg/test-quik-tech/cmd/config/connections"
	"github.com/sbuttigieg/test-quik-tech/cmd/config/store"
	"github.com/sbuttigieg/test-quik-tech/cmd/config/wallet/api"
	"github.com/sbuttigieg/test-quik-tech/wallet/http/api/middleware"
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func main() {
	// env variables
	endpointURL := os.Getenv("ENDPOINT_URL")
	apiAddr := os.Getenv("API_PORT")

	ctx := context.Background()

	// config
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	// connections
	redisConnection, err := connections.NewRedis()
	if err != nil {
		log.Fatal(err.Error())
	}

	// redis setup
	cache, err := store.NewCache(c, redisConnection)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	User1 := models.User{
		WalletID:     "6cc4ee0d-9919-4857-a70d-9b7283957e16",
		Balance:      100,
		Username:     "Bob",
		Password:     "123456",
		LastActivity: time.Now(),
	}

	User2 := models.User{
		WalletID:     "0924f01f-3f70-4fe4-ac82-dce4b30e2a7f",
		Balance:      100,
		Username:     "Joe",
		Password:     "654321",
		LastActivity: time.Now(),
	}

	User3 := models.User{
		WalletID:     "d2ba410a-9bc4-476b-86af-c55525b527df",
		Balance:      100,
		Username:     "Dave",
		Password:     "456789",
		LastActivity: time.Now(),
	}

	err = cache.SetKey(User1.WalletID, User1, c.CacheExpiry)
	if err != nil {
		fmt.Println("SetKey", err)
	}

	err = cache.SetKey(User2.WalletID, User2, c.CacheExpiry)
	if err != nil {
		fmt.Println("SetKey", err)
	}

	err = cache.SetKey(User3.WalletID, User3, c.CacheExpiry)
	if err != nil {
		fmt.Println("SetKey", err)
	}

	u1, ok := cache.GetKeyBytes(User1.WalletID)
	if !ok {
		fmt.Println("GetKey User1", ok)
	}

	u2, ok := cache.GetKeyBytes(User2.WalletID)
	if !ok {
		fmt.Println("GetKey User1", ok)
	}

	u3, ok := cache.GetKeyBytes(User3.WalletID)
	if !ok {
		fmt.Println("GetKey User1", ok)
	}

	var d1, d2, d3 models.User

	err = json.Unmarshal(u1, &d1)
	if err != nil {
		fmt.Println("Unmarshal d1", err)
	}

	err = json.Unmarshal(u2, &d2)
	if err != nil {
		fmt.Println("Unmarshal d2", err)
	}

	err = json.Unmarshal(u3, &d3)
	if err != nil {
		fmt.Println("Unmarshal d3", err)
	}

	fmt.Println("User1 ==>", d1)
	fmt.Println("User2 ==>", d2)
	fmt.Println("User3 ==>", d3)

	// api setup
	apiService, err := api.NewService(c, cache, uuid.New, time.Now)
	if err != nil {
		log.Fatal(err.Error())
	}

	apiHandlers, err := api.NewHandlers(apiService)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Comment for debug mode. Uncomment for production
	// gin.SetMode(gin.ReleaseMode)

	// Create a new instance of the Gin router
	apiRouter := gin.New()

	err = apiRouter.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Endpoints
	apiRouter.POST(fmt.Sprintf("%s/:wallet_id/auth", endpointURL), apiHandlers.Auth)
	apiRouter.GET(fmt.Sprintf("%s/:wallet_id/balance", endpointURL), middleware.BasicAuth(), apiHandlers.Balance)
	apiRouter.POST(fmt.Sprintf("%s/:wallet_id/credit", endpointURL), middleware.BasicAuth(), apiHandlers.Credit)
	apiRouter.POST(fmt.Sprintf("%s/:wallet_id/debit", endpointURL), middleware.BasicAuth(), apiHandlers.Debit)

	// Start the server
	err = apiRouter.Run(fmt.Sprintf(":%s", apiAddr))

	if err != nil {
		log.Fatal(err.Error())
	}
}
