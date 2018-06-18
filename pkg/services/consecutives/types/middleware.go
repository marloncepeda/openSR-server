package types

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

func createType(data template, db *gorm.DB) (int, string) {

	isAvailable := availability(data.Name, db)

	if !isAvailable {
		return http.StatusSeeOther, "The type is been used by another entity"
	}

	newType := Model{}

	id, err := newType.last(db)

	if err != nil {
		id = "0"
	}

	pk, _ := strconv.Atoi(decrypt(id))

	newType.ID = encrypt(strconv.Itoa(pk + 1))
	newType.Name = encrypt(data.Name)
	newType.Description = encrypt(data.Description)

	err = newType.save(db)

	if err != nil {
		return http.StatusInternalServerError, "Error creating the consecutive type"
	}

	return http.StatusCreated, "type has been created"
}

func getTypes(db *gorm.DB) ([]Model, error) {

	types := Model{}

	result, err := types.types(db)

	if err != nil {
		return nil, err
	}

	return decodes(result), nil
}

func getType(id string, db *gorm.DB) (Model, error) {

	types := Model{}

	result, err := types.query("id", encrypt(id), db)

	if err != nil {
		return result, err
	}

	return decode(result), nil
}
