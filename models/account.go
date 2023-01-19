package models

import (
	u "digit-liblary/utils"
	"strings"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email string `json:"email"`
  Password string `json:"password"`
  Token string `json:"token"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is not valid"), false
	}

  if len(account.Password) < 6 {
    return u.Message(false, "Password is required"), false
  }

	temp := &Account{}

	err := GetDataBase().
}