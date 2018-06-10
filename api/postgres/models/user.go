package models

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

// Save function validate the User struct, and save a new User on the database
func (u *User) Save(db *gorm.DB) error {

	err := db.Create(u).Error

	if err != nil {
		return err
	}

	return nil
}

// UpdatePassword function ....
func (u *User) UpdatePassword(db *gorm.DB) error {

	err := db.Model(u).Update("password", u.Password).Error

	if err != nil {
		return err
	}

	return nil
}

// UpdateFields function ....
func (u *User) UpdateFields(db *gorm.DB) error {

	err := db.Model(u).Updates(u).Error

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

// CheckInformation ....
func (u *User) CheckInformation(db *gorm.DB) error {

	user := User{}
	user.ID = u.ID

	err := db.First(&user).Error

	if err != nil {
		return err
	}

	return nil

}

// GetUsers ...
func (u *User) GetUsers(db *gorm.DB) ([]User, error) {

	var users []User

	err := db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser ...
func (u *User) GetUser(db *gorm.DB) (User, error) {

	var user User

	err := db.Where("ID = ?", u.ID).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil

}
