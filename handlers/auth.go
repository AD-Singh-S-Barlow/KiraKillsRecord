
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
    "github.com/AD-Singh-S-Barlow/KiraKillsRecord/models"

	"github.com/dgrijalva/jwt-go"
	
	"time"

)

func Register(c *gin.Context) {
	var account models.Account
// Bind the JSON payload to the 'account' variable
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
// Append the 'account' to the 'models.Accounts' slice
	models.Accounts = append(models.Accounts, account)

	c.JSON(http.StatusOK, gin.H{
		"message": "Account registered successfully",
		"account": account,
	})
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
// Bind the JSON payload to the 'loginData' variable
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, account := range models.Accounts {
		if account.Username == loginData.Username && account.Password == loginData.Password {
			// Create the JWT token
			token := createToken(account.Username)

			// Authentication successful, include the token in the response
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"token":   token,
			})
			return
		}
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

func createToken(username string) string {
	// Define the expiration time for the token (e.g., 1 hour from now)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims, including the username and expiration time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generateing the token using the claims and a signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using a secret key
	secret := []byte("36e9307fac0cce1e4509e2d83c8eb518d4abf1d37cb84d632fb8a40e84fa5542")
	tokenString, _ := token.SignedString(secret)

	return tokenString
}
