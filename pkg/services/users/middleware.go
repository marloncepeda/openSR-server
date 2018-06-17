package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func createUser(data template, db *gorm.DB) (int, string) {

	isAvailable := availability(data.Username, db)

	if !isAvailable {
		return http.StatusSeeOther, "The username is being used"
	}

	//instance
	user := Model{}

	//select the PK
	id, err := user.last(db)

	if err != nil {
		id = "0"
	}

	pk, _ := strconv.Atoi(decrypt(id))
	uuid, _ := uuid.NewV4()

	user.ID = encrypt(strconv.Itoa(pk + 1))
	user.PublicID = encrypt(uuid.String())
	user.Name = encrypt(data.Name)
	user.Surname = encrypt(data.Surname)
	user.SecondSurName = encrypt(data.SecondSurname)
	user.Phone = encrypt(data.Phone)
	user.UserName = encrypt(data.Username)
	user.Password = encrypt(data.Password)

	err = user.save(db)

	if err != nil {
		return http.StatusInternalServerError, "Error creating the user"
	}
	return http.StatusCreated, "user has been created"
}

func getAllUsers(db *gorm.DB) ([]users, error) {

	user := Model{}

	users, err := user.users(db)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func deleteUser(public string, db *gorm.DB) (int, string) {

	user := Model{PublicID: encrypt(public)}

	_, err := user.query("public_id", encrypt(public), db)

	if err != nil {
		return http.StatusNotFound, "User does not exists on the database"
	}

	err = user.delete(db)

	if err != nil {
		return http.StatusInternalServerError, "Error deleting the user"
	}

	return http.StatusOK, fmt.Sprintf("User %s deleted", public)
}

func updateUser(data map[string]string, db *gorm.DB) (int, string) {

	field, value, id := data["field"], data["value"], data["id"]

	variants := []string{"name", "surname", "phone", "password"}

	var isValid bool

	for _, parameter := range variants {

		if parameter == field {
			isValid = true
		}
	}

	isValid = false

	if !isValid {
		return http.StatusBadRequest, "This field is not allowed in this API"
	}

	user := Model{}
	user.PublicID = encrypt(id)

	_, err := user.query("public_id", user.PublicID, db)

	if err != nil {
		return http.StatusNotFound, "User does not exists on the database"
	}

	switch field {

	case "name":
		user.Name = encrypt(value)

	case "surname":
		user.Surname = encrypt(value)

	case "phone":
		user.Phone = encrypt(value)

	case "password":
		user.Password = encrypt(value)

	}

	err = user.update(db)

	if err != nil {
		return http.StatusInternalServerError, "Error deleting the user"
	}
	return http.StatusOK, fmt.Sprintf("User %s updated", id)
}
