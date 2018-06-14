package user

import (
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	ID            string `gorm:"primary_key"`
	Name          string `gorm:"not null"`
	Surname       string `gorm:"not null"`
	SecondSurName string `gorm:"not null"`
	Phone         string `gorm:"not null"`
	UserName      string `gorm:"unique; not null; index:username"`
	Password      string `gorm:"not null"`
}

// TableName return a table name
func (User) TableName() string { return "Users" }

func (u *User) save(db *gorm.DB) error {

	err := db.Create(u).Error

	if err != nil {
		return err
	}

	return nil
}

// Update function ....
func (u *User) update(db *gorm.DB) error {

	err := db.Model(User{}).Updates(u).Error

	if err != nil {
		return err
	}

	return nil
}

// Delete ...
func (u *User) Delete(db *gorm.DB) error {

	err := db.Delete(u).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *User) username(db *gorm.DB) (string, error) {

	var user User

	err := db.Where("user_name = ?", u.UserName).First(&user).Error

	if err != nil {
		return "", err
	}

	return user.UserName, nil
}

func (u *User) users(db *gorm.DB) ([]User, error) {

	var users []User

	err := db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) user(db *gorm.DB) (User, error) {

	var user User

	err := db.Where("ID = ?", u.ID).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) count(db *gorm.DB) (string, error) {

	var user User

	err := db.Select("name, ID").Find(&user).Error

	if err != nil {
		return "", err
	}

	return user.ID, nil
}

// Login ...
func (u *User) Login(db *gorm.DB) (bool, string) {

	var user User

	err := db.Where("user_name = ? AND password = ?", u.UserName, u.Password).First(&user).Error

	if err != nil {
		return false, user.ID
	}

	return true, user.ID
}
