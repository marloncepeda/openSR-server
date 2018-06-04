package user

import (
	"fmt"

	"github.com/ctreminiom/scientific-logs-api/api/postgres/models"
	"github.com/ctreminiom/scientific-logs-api/api/security/aes"
	"github.com/go-pg/pg"
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

func check(username string, db *pg.DB) bool {

	usernameEncoded := encrypt(username)

	fmt.Println(usernameEncoded)

	user := new(models.User)

	err := db.Model(user).Column("username").Where("username = ?", usernameEncoded).Select()

	fmt.Println(err)

	if err != nil {
		return false
	}

	return true
}

func encrypt(text string) string { return aes.Encrypt(text) }

func decrypt(text string) string { return aes.Decrypt(text) }
