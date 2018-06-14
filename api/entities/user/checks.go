package user

import (
	"github.com/jinzhu/gorm"
)

type creation struct {
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

func availability(userName string, db *gorm.DB) bool {

	user := User{UserName: userName}

	queryResult, _ := user.username(db)

	if userName == queryResult {
		return false
	}

	return true
}
