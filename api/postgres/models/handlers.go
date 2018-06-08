package models

import (
	"errors"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Migrate ...
func Migrate(db *pg.DB) error {

	options := &orm.CreateTableOptions{IfNotExists: true}

	err := db.CreateTable(&User{}, options)

	if err != nil {
		return errors.New("Error while creating tables" + err.Error())
	}

	return nil

}
