package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ctreminiom/openSR-server/api/security"
	"github.com/jinzhu/gorm"
)

func create(body creation, db *gorm.DB) (HTTPCode int, response string) {

	isAvailable := availability(body.Username, db)

	if !isAvailable {
		return http.StatusBadRequest, "the username is being used"
	}

	user := User{}

	nextID, err := user.count(db)

	if err != nil {
		return http.StatusInternalServerError, "error feching the user count"
	}

	primaryKey, _ := strconv.Atoi(nextID)

	user.ID = encrypt(strconv.Itoa(primaryKey + 1))
	user.Name = encrypt(body.Name)
	user.Surname = encrypt(body.Surname)
	user.SecondSurName = encrypt(body.SecondSurname)
	user.Phone = encrypt(body.Phone)
	user.UserName = encrypt(body.Username)
	user.Password = encrypt(body.Password)

	err = user.save(db)

	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, "user has been created"
}

func erase(id string, db *gorm.DB) (HTTPCode int, response string) {

	user := User{ID: encrypt(id)}

	_, err := user.user(db)

	if err != nil {
		return http.StatusNotFound, "User does not exist on the database"
	}

	user.Delete(db)

	return http.StatusOK, fmt.Sprintf("User %s deleted", id)
}

func update(id string, db *gorm.DB) (HTTPCode int, response string) {

	user := User{ID: encrypt(id)}

	_, err := user.user(db)

	if err != nil {
		return http.StatusNotFound, "User does not exist on the database"
	}

	user.update(db)

	return http.StatusOK, fmt.Sprintf("User %s updated", id)
}

func getUsers(db *gorm.DB) (HTTPCode int, users *[]User, responde string) {

	user := User{}

	data, err := user.users(db)

	if err != nil {
		return http.StatusInternalServerError, nil, "error feching all users"
	}

	return http.StatusOK, decodes(data), ""

}

func prue(id string, db *gorm.DB) (HTTPCode int, u User, response string) {

	user := User{ID: encrypt(id)}

	//data, err := user.user(db)

	return 1, user, ""

}

func getUser(db *gorm.DB) (HTTPCode int, users User, responde string) {

	user := User{}

	errUser := User{ID: "error"}

	data, err := user.user(db)

	if err != nil {
		return http.StatusInternalServerError, errUser, "error feching all users"
	}

	return http.StatusOK, data, ""
}

func decodes(u []User) *[]User {

	var users []User

	for _, row := range u {

		users = append(users, decode(row))

	}

	return &users
}

func decode(u User) User {

	var user User

	// ADD CONCURRENCY ***********
	user.ID = decrypt(u.ID)
	user.Name = decrypt(u.Name)
	user.Surname = decrypt(u.Surname)
	user.SecondSurName = decrypt(u.SecondSurName)
	user.Phone = decrypt(u.Phone)
	user.UserName = decrypt(u.UserName)
	user.Password = decrypt(u.Password)
	// ADD CONCURRENCY ***********

	return user
}

func encrypt(text string) string { return security.Encrypt(text) }

func decrypt(text string) string { return security.Decrypt(text) }
