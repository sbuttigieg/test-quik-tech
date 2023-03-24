package main

import (
	"context"
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

	// logger setup
	logFile := "logs.txt"

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	log := config.NewLogger(c, f)

	// connections
	redisConnection, err := connections.NewRedis()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbConnection, err := connections.NewMySQL(c, log)
	if err != nil {
		log.Fatal(err.Error())
	}

	// store setup
	err = store.DBInit(dbConnection)
	if err != nil {
		log.Error("Database Initiation Error: ", err)
	}

	// redis setup
	cache := store.NewCache(c, redisConnection)

	// api setup
	apiStore := api.NewStore(dbConnection)
	apiService := api.NewService(c, cache, apiStore, log, uuid.New, time.Now)
	apiHandlers := api.NewHandlers(apiService)

	// Comment for debug mode. Uncomment for production
	// gin.SetMode(gin.ReleaseMode)

	// Create a new instance of the Gin router
	apiRouter := gin.New()
	apiRouter.Use(gin.Recovery())
	apiRouter.Use(middleware.Logger(ctx, log))

	err = apiRouter.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Endpoints
	apiRouter.POST(fmt.Sprintf("%s/:wallet_id/auth", endpointURL), apiHandlers.Auth)
	apiRouter.GET(fmt.Sprintf("%s/:wallet_id/balance", endpointURL), middleware.BasicAuth(apiService), apiHandlers.Balance)
	apiRouter.POST(fmt.Sprintf("%s/:wallet_id/credit", endpointURL), middleware.BasicAuth(apiService), apiHandlers.Credit)
	apiRouter.POST(fmt.Sprintf("%s/:wallet_id/debit", endpointURL), middleware.BasicAuth(apiService), apiHandlers.Debit)

	// Start the server
	err = apiRouter.Run(fmt.Sprintf(":%s", apiAddr))

	if err != nil {
		log.Fatal(err.Error())
	}
}
