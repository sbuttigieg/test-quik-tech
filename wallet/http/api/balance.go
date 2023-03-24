package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func (h *Handler) Balance(c *gin.Context) {
	balance, err := h.service.Balance(c.Param("wallet_id"))
	if err != nil {
		c.Abort()

		switch err.Error() {
		case PlayerNotFoundError:
			c.JSON(http.StatusBadRequest, "Player does not exist")
		case PlayerNotLoggedIn:
			c.JSON(http.StatusBadRequest, "Player not logged in")
		default:
			c.JSON(http.StatusInternalServerError, "Error retrieving balance")
		}

		return
	}

	c.JSON(http.StatusOK, models.Balance{
		WalletID: c.Param("wallet_id"),
		Balance:  *balance,
	})
}
