package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/sbuttigieg/test-quik-tech/cmd/config/app/api"
	"github.com/sbuttigieg/test-quik-tech/internal/http/api/middleware"
)

func main() {
	endpointURL := os.Getenv("ENDPOINT_URL")
	apiAddr := os.Getenv("API_PORT")

	apiService, err := api.NewService(uuid.New, time.Now)
	if err != nil {
		log.Fatal(err.Error())
	}

	apiHandlers, err := api.NewHandlers(apiService)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Cooment for debug mode. Uncomment for production
	// gin.SetMode(gin.ReleaseMode)

	// Create a new instance of the Gin router
	apiRouter := gin.New()
	apiRouter.SetTrustedProxies(nil)

	// Endpoints
	apiRouter.POST(fmt.Sprintf("%s/:wallet_id/auth", endpointURL), apiHandlers.Auth)
	apiRouter.GET(fmt.Sprintf("%s/:wallet_id/balance", endpointURL), middleware.BasicAuth(), apiHandlers.Balance)
	apiRouter.POST(fmt.Sprintf("%s/:wallet_id/credit", endpointURL), middleware.BasicAuth(), apiHandlers.Credit)
	apiRouter.POST(fmt.Sprintf("%s/:wallet_id/debit", endpointURL), middleware.BasicAuth(), apiHandlers.Debit)

	// Start the server
	apiRouter.Run(fmt.Sprintf(":%s", apiAddr))
}
