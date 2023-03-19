package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-quik-tech/internal/models"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Auth(c *gin.Context) {
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

	var req AuthRequest

	if err := c.BindJSON(&req); err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, "Error reading request")
		return
	}

	if req.Username == "" || req.Password == "" {
		c.Abort()
		c.JSON(http.StatusBadRequest, "Missing Username and/or Password")
		return
	}

	var valid bool

	switch c.Param("wallet_id") {
	case User1.WalletID:
		valid = CheckCredentials(User1, req)
	case User2.WalletID:
		valid = CheckCredentials(User2, req)
	case User3.WalletID:
		valid = CheckCredentials(User3, req)
	}

	if !valid {
		c.Abort()
		c.JSON(http.StatusUnauthorized, "Incorrect Username and/or Password")
		return
	}

	// **** To set User as Logged-in in cache ****

	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func CheckCredentials(user models.User, req AuthRequest) bool {
	if user.Username != req.Username || user.Password != req.Password {
		return false
	}

	return true
}
