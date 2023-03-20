package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

type CreditRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

func (h *Handler) Credit(c *gin.Context) {
	var req CreditRequest

	if err := c.BindJSON(&req); err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, "Error reading request")

		return
	}

	balance, err := h.service.Credit(c.Param("wallet_id"), req.Amount)
	if err != nil {
		c.Abort()

		switch err.Error() {
		case "player not found":
			c.JSON(http.StatusBadRequest, "Player does not exist")
		case "negative value":
			c.JSON(http.StatusBadRequest, "Negative value error")
		default:
			c.JSON(http.StatusInternalServerError, "Error processing credit transaction")
		}

		return
	}

	c.JSON(http.StatusOK, models.Transaction{
		WalletID: c.Param("wallet_id"),
		Amount:   req.Amount,
		Type:     req.Description,
		Balance:  balance,
	})
}
