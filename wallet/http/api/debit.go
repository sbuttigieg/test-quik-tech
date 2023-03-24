package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type DebitRequest struct {
	Amount      decimal.Decimal `json:"amount"`
	Description string          `json:"description"`
}

func (h *Handler) Debit(c *gin.Context) {
	var req DebitRequest

	if err := c.BindJSON(&req); err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, "Error reading request")

		return
	}

	transaction, err := h.service.Debit(c.Param("wallet_id"), req.Description, req.Amount)
	if err != nil {
		c.Abort()

		switch err.Error() {
		case PlayerNotFoundError:
			c.JSON(http.StatusBadRequest, "Player does not exist")
		case NegativeValueError:
			c.JSON(http.StatusBadRequest, "Negative value error")
		case InsufficientFundsError:
			c.JSON(http.StatusBadRequest, "Insufficient  Funds")
		case PlayerNotLoggedIn:
			c.JSON(http.StatusBadRequest, "Player not logged in")
		default:
			c.JSON(http.StatusInternalServerError, "Error processing debit transaction")
		}

		return
	}

	c.JSON(http.StatusOK, transaction)
}
