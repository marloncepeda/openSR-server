package user

import (
	"github.com/ctreminiom/scientific-logs-api/api/security/aes"
)

type registerTemplate struct {
	Name          string `form:"name" json:"name" binding:"required"`
	Surname       string `form:"surname" json:"surname" binding:"required"`
	SecondSurname string `form:"secondSurname" json:"secondSurname" binding:"required"`
	Phone         string `form:"phone" json:"phone" binding:"required"`
	Username      string `form:"username" json:"username" binding:"required"`
	Password      string `form:"password" json:"password" binding:"required"`
}

func validate(json error) bool {

	if json != nil {
		return false
	}

	return true
}

func encrypt(text string) string { return aes.Encrypt(text) }
