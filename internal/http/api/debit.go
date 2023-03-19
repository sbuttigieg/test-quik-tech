package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-quik-tech/internal/models"
)

type DebitRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

func (h *Handler) Debit(c *gin.Context) {
	var req DebitRequest

	if err := c.BindJSON(&req); err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, "Error reading request")
		return
	}

	if req.Amount < 0 {
		c.Abort()
		c.JSON(http.StatusBadRequest, "Negative value error")
		return
	}

	// **** To get balance from cache ****
	// **** To deduct debit amount from balance ****
	// **** Throw error if balance will be less than 0 after transaction ****
	// **** To update new balance in store and cache ****

	c.JSON(http.StatusOK, models.Transaction{
		WalletID: c.Param("wallet_id"),
		Amount:   req.Amount,
		Type:     req.Description,
		Balance:  100,
	})
}
