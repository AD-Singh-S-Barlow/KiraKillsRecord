// handlers/auth.go
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
	// Handle login logic here
}
