package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// **** To be replaced by users from DB or cache ****
		Player1 := models.Player{
			WalletID: "6cc4ee0d-9919-4857-a70d-9b7283957e16",
			Username: "Bob",
			Password: "123456",
		}

		Player2 := models.Player{
			WalletID: "0924f01f-3f70-4fe4-ac82-dce4b30e2a7f",
			Username: "Joe",
			Password: "654321",
		}

		Player3 := models.Player{
			WalletID: "d2ba410a-9bc4-476b-86af-c55525b527df",
			Username: "Dave",
			Password: "456789",
		}
		// **************************************************

		var username, password string

		switch c.Param("wallet_id") {
		case Player1.WalletID:
			username = Player1.Username
			password = Player1.Password
		case Player2.WalletID:
			username = Player2.Username
			password = Player2.Password
		case Player3.WalletID:
			username = Player3.Username
			password = Player3.Password
		}

		un, pwd, hasAuth := c.Request.BasicAuth()

		if !hasAuth || un != username || pwd != password {
			c.Abort()
			c.JSON(http.StatusUnauthorized, "Authentication Error")

			return
		}
	}
}
