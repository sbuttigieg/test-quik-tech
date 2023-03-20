package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Auth(c *gin.Context) {
	var req AuthRequest

	if err := c.BindJSON(&req); err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, "Error reading request")

		return
	}

	user, err := h.service.Auth(c.Param("wallet_id"), req.Username, req.Password)
	if err != nil {
		c.Abort()

		switch err.Error() {
		case "user not found":
			c.JSON(http.StatusBadRequest, "User does not exist")
		case "missing credentials":
			c.JSON(http.StatusBadRequest, "Missing Username and/or Password")
		case "incorrect credentials":
			c.JSON(http.StatusUnauthorized, "Incorrect Username and/or Password")
		default:
			c.JSON(http.StatusInternalServerError, "Error processing user authentication")
		}

		return
	}

	c.JSON(http.StatusOK, models.Balance{
		WalletID: c.Param("wallet_id"),
		Balance:  user.Balance,
	})
}
