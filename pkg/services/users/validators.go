package users

import (
	"github.com/ctreminiom/openSR-server/pkg/auth"
	"github.com/jinzhu/gorm"
)

type template struct {
	Name          string `form:"name" json:"name" binding:"required"`
	Surname       string `form:"surname" json:"surname" binding:"required"`
	SecondSurname string `form:"secondSurname" json:"secondSurname" binding:"required"`
	Phone         string `form:"phone" json:"phone" binding:"required"`
	Username      string `form:"username" json:"username" binding:"required"`
	Password      string `form:"password" json:"password" binding:"required"`
}

func check(json error) bool {

	if json != nil {
		return false
	}

	return true
}

func availability(nickName string, db *gorm.DB) bool {

	user := Model{}

	query, _ := user.query("user_name", encrypt(nickName), db)

	if encrypt(nickName) == query.UserName {
		return false
	}

	return true

}

func decodes(u []users) []users {

	var result []users

	for _, row := range u {

		result = append(result, decode(row))
	}

	return result
}

func decode(u users) users {

	result := users{}

	result.PublicID = decrypt(u.PublicID)
	result.Name = decrypt(u.Name)
	result.Surname = decrypt(u.Surname)
	result.SecondSurName = decrypt(u.SecondSurName)
	result.Phone = decrypt(u.Phone)
	result.UserName = decrypt(u.UserName)

	return result
}

func encrypt(text string) string { return auth.Encrypt(text) }

func decrypt(text string) string { return auth.Decrypt(text) }
