package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-quik-tech/internal/models"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// **** To be replaced by users from DB or cache ****
		User1 := models.User{
			WalletID: "6cc4ee0d-9919-4857-a70d-9b7283957e16",
			Username: "Bob",
			Password: "123456",
		}

		User2 := models.User{
			WalletID: "0924f01f-3f70-4fe4-ac82-dce4b30e2a7f",
			Username: "Joe",
			Password: "654321",
		}

		User3 := models.User{
			WalletID: "d2ba410a-9bc4-476b-86af-c55525b527df",
			Username: "Dave",
			Password: "456789",
		}
		// **************************************************

		var username, password string

		switch c.Param("wallet_id") {
		case User1.WalletID:
			username = User1.Username
			password = User1.Password
		case User2.WalletID:
			username = User2.Username
			password = User2.Password
		case User3.WalletID:
			username = User3.Username
			password = User3.Password
		}

		user, pwd, hasAuth := c.Request.BasicAuth()

		if !hasAuth || user != username || pwd != password {
			c.Abort()
			c.JSON(http.StatusUnauthorized, "Authentication Error")
			return
		}
	}
}
