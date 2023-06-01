// models/accounts.go
package models

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var Accounts []Account
