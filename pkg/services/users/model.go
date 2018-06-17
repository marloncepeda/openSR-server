package users

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Model ...
type Model struct {
	ID            string `gorm:"primary_key"`
	PublicID      string `gorm:"not null"`
	Name          string `gorm:"not null"`
	Surname       string `gorm:"not null"`
	SecondSurName string `gorm:"not null"`
	Phone         string `gorm:"not null"`
	UserName      string `gorm:"unique; not null"`
	Password      string `gorm:"not null"`
}

// TableName .....
func (Model) TableName() string { return "users" }

func (u *Model) save(db *gorm.DB) error {

	err := db.Create(u).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *Model) update(db *gorm.DB) error {

	err := db.Model(Model{}).Omit("ID", "public_id").Updates(u).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *Model) delete(db *gorm.DB) error {

	user := Model{}

	// Search the PK of the user using the public_id
	err := db.Where("public_id = ?", u.PublicID).First(&user).Error

	if err != nil {
		return err
	}

	// Delete the user using the PK --> GORM restriccion
	err = db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *Model) last(db *gorm.DB) (string, error) {

	user := Model{}

	err := db.Last(&user).Error

	if err != nil {
		return "0", err
	}

	return user.ID, nil
}

type users struct {
	PublicID      string
	Name          string
	Surname       string
	SecondSurName string
	Phone         string
	UserName      string
}

func (u *Model) users(db *gorm.DB) ([]users, error) {

	result := []users{}

	err := db.Select("public_id, name, surname, second_sur_name, phone, user_name").Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (u *Model) query(where, value string, db *gorm.DB) (users, error) {

	user := users{}

	var err error

	switch where {

	case "user_name":
		err = db.Where("user_name = ?", value).Select("public_id, name, surname, second_sur_name, phone, user_name").First(&user).Error

	case "name":
		err = db.Where("name = ?", value).Select("public_id, name, surname, second_sur_name, phone, user_name").First(&user).Error

	case "surname":
		err = db.Where("surname = ?", value).Select("public_id, name, surname, second_sur_name, phone, user_name").First(&user).Error

	case "phone":
		err = db.Where("phone = ?", value).Select("public_id, name, surname, second_sur_name, phone, user_name").First(&user).Error

	case "public_id":
		err = db.Where("public_id = ?", value).Select("public_id, name, surname, second_sur_name, phone, user_name").First(&user).Error

	default:
		err = errors.New("Invalid parameter")

	}

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *Model) login(db *gorm.DB) (bool, string) {

	var user Model

	err := db.Where("user_name = ? AND password = ?", u.UserName, u.Password).First(&user).Error

	if err != nil {
		return false, user.ID
	}

	return true, user.PublicID
}
