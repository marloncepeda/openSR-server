package user

import (
	"net/http"

	"github.com/ctreminiom/openSR-server/api/postgres/models"
	"github.com/ctreminiom/openSR-server/api/security"

	"github.com/jinzhu/gorm"
)

func createUser(body registerJSON, db *gorm.DB) (httpCode int, message string) {

	isBeingUsed := checkTheUsernameAvailability(body.Username, db)

	if isBeingUsed {
		return http.StatusBadRequest, "the user is being used"
	}

	newUser := new(models.User)

	newUser.ID = "sadsad"
	newUser.Name = body.Name
	newUser.Surname = body.Surname
	newUser.SecondSurName = body.SecondSurname
	newUser.Phone = body.Phone
	newUser.UserName = body.Username
	newUser.Password = body.Password

	err := newUser.Save(db)

	if err != nil {
		return http.StatusBadRequest, "LOL"
	}

	return http.StatusOK, "user created"

	/*

		var count int

		db.Model(user).Count(count)

		fmt.Println(count)

		return http.StatusOK, "user created"
	*/
}

func encrypt(text string) string { return security.Encrypt(text) }

func decrypt(text string) string { return security.Decrypt(text) }
