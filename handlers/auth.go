
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
    "github.com/AD-Singh-S-Barlow/KiraKillsRecord/models"

)

func Register(c *gin.Context) {
	var account models.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Accounts = append(models.Accounts, account)

	c.JSON(http.StatusOK, gin.H{
		"message": "Account registered successfully",
		"account": account,
	})
}

func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, account := range models.Accounts {
		if account.Username == loginData.Username && account.Password == loginData.Password {
			// Authentication successful
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
			return
		}
	}

	// Authentication failed
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
}

