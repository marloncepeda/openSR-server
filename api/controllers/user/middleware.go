package user

import (
	"net/http"
	"strconv"

	"github.com/ctreminiom/scientific-logs-api/api/postgres/models"
	base64 "github.com/ctreminiom/scientific-logs-api/api/security/base64"
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

func validate(json error) bool {

	if json != nil {
		return false
	}

	return true
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

func create(body registerTemplate, db *pg.DB) (code int, message string) {

	used := check(body.Username, db)

	if used {
		return http.StatusBadRequest, "username already exists"
	}

	//Create user
	count, _ := db.Model((*models.User)(nil)).Count()

	newUser := models.User{
		ID:            encrypt(strconv.Itoa(count + 1)),
		Consecutive:   encrypt("i"),
		Name:          encrypt(body.Name),
		Surname:       encrypt(body.Surname),
		SecondSurName: encrypt(body.SecondSurname),
		Phone:         encrypt(body.Phone),
		Username:      encrypt(body.Username),
		Password:      encrypt(body.Password),
	}

	err := newUser.Save(db)

	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	return http.StatusOK, "user created"

}

func encrypt(text string) string { return base64.Encrypt(text) }

func decrypt(text string) string { return base64.Decrypt(text) }
