package models

import (
	"errors"
	"log"
	"time"

	"github.com/go-pg/pg"
)

// User ....
type User struct {
	tableName struct{} `sql:"users"`

	ID            string `sql:"type:text, pk"`
	Consecutive   string `sql:"type:text"`
	Name          string `sql:"type:text"`
	Surname       string `sql:"type:text"`
	SecondSurName string `sql:"type:text"`
	Phone         string `sql:"type:text"`
	Username      string `sql:"type:text, unique"`
	Password      string `sql:"type:text"`

	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time
}

// Save saved a User
func (user *User) Save(db *pg.DB) error {

	err := db.Insert(user)

	if err != nil {
		return errors.New("Error while inserting the new User " + err.Error())
	}

	log.Println("User" + user.Name + "Created")

	return nil

}
