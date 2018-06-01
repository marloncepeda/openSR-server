package models

import (
	"errors"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Migrate ...
func Migrate(db *pg.DB) error {

	options := &orm.CreateTableOptions{IfNotExists: true}

	err := db.CreateTable(&User{}, options)

	if err != nil {
		log.Panic(errors.New("Error while creating tables" + err.Error()))
		return errors.New("Error while creating tables" + err.Error())
	}

	log.Printf("SRIV")
	return nil

}
