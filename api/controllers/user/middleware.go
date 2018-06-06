package user

import (
	"net/http"
	"strconv"

	"github.com/ctreminiom/openSR-server/api/postgres/models"
	base64 "github.com/ctreminiom/openSR-server/api/security/base64"
	pg "github.com/go-pg/pg"
)

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
