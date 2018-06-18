package types

import (
	"github.com/ctreminiom/openSR-server/pkg/auth"
	"github.com/jinzhu/gorm"
)

type template struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

func check(json error) bool {

	if json != nil {
		return false
	}

	return true
}

func availability(name string, db *gorm.DB) bool {

	consecutiveType := Model{}

	query, _ := consecutiveType.query("name", encrypt(name), db)

	if encrypt(name) == query.Name {
		return false
	}

	return true

}

func decodes(t []Model) []Model {

	var result []Model

	for _, row := range t {

		result = append(result, decode(row))
	}

	return result
}

func decode(t Model) Model {

	result := Model{}

	result.ID = decrypt(t.ID)
	result.Name = decrypt(t.Name)
	result.Description = decrypt(t.Description)

	return result
}

func encrypt(text string) string { return auth.Encrypt(text) }

func decrypt(text string) string { return auth.Decrypt(text) }
