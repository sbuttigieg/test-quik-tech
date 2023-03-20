package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sbuttigieg/test-quik-tech/wallet/services/api"
)

func BasicAuth(s api.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, hasAuth := c.Request.BasicAuth()

		if !hasAuth {
			c.Abort()
			c.JSON(http.StatusUnauthorized, "Missing Authentication")

			return
		}

		_, err := s.Auth(c.Param("wallet_id"), username, password)
		if err != nil {
			c.Abort()

			switch err.Error() {
			case "player not found":
				c.JSON(http.StatusBadRequest, "Player does not exist")
			default:
				c.JSON(http.StatusUnauthorized, "Authentication Error")
			}

			return
		}
	}
}
