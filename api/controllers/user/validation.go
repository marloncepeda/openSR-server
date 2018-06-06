package user

import (
	"github.com/ctreminiom/openSR-server/api/postgres/models"
	pg "github.com/go-pg/pg"
)

type registerTemplate struct {
	Name          string `form:"name" json:"name" binding:"required"`
	Surname       string `form:"surname" json:"surname" binding:"required"`
	SecondSurname string `form:"secondSurname" json:"secondSurname" binding:"required"`
	Phone         string `form:"phone" json:"phone" binding:"required"`
	Username      string `form:"username" json:"username" binding:"required"`
	Password      string `form:"password" json:"password" binding:"required"`
}

type loginTemplate struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func validateHTTPBody(json error) bool {

	if json != nil {
		return false
	}

	return true
}

func confirmUsername(username, password string, db *pg.DB) (bool, error) {

	user := new(models.User)

	err := db.Model(user).Column("username").Where("username = ?", encrypt(username)).Where("password = ?", encrypt(password)).Select()

	if err != nil {
		return false, err
	}

	return true, nil
}

func check(username string, db *pg.DB) bool {

	usernameEncoded := encrypt(username)

	user := new(models.User)

	err := db.Model(user).Column("username").Where("username = ?", usernameEncoded).Select()

	if err != nil {
		return false
	}

	return true
}
