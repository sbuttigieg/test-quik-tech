package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-quik-tech/internal/models"
)

func (h *Handler) Balance(c *gin.Context) {
	// **** To get balance from cache ****
	c.JSON(http.StatusOK, models.Balance{
		WalletID: c.Param("wallet_id"),
		Balance:  100,
	})
}
